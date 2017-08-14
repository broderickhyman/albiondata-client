package client

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	photon "github.com/hmadison/photon_spectator"
	"github.com/regner/albionmarket-client/log"
)

type listener struct {
	handle     *pcap.Handle
	source     *gopacket.PacketSource
	deviceName string
	fragments  *photon.FragmentBuffer
	running    bool
	stopping   bool
	quit       chan bool
	router     *Router
}

func newListener(router *Router) *listener {
	return &listener{
		fragments: photon.NewFragmentBuffer(),
		running:   false,
		stopping:  false,
		quit:      make(chan bool),
		router:    router,
	}
}

func (l *listener) createOnline(deviceName string) {
	log.Debugf("Creating listener %s...", deviceName)
	l.deviceName = deviceName

	handle, err := pcap.OpenLive(deviceName, 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	l.handle = handle
}

func (l *listener) setPortsAndRun(deviceName string, ports []int) {
	// First quit itself so its not parsing garbage
	if l.running {
		l.quit <- true
	}

	log.Debugf("Setting ports %v on device %s", ports, deviceName)

	// Now create a BPF Filter string
	filter := ""
	for _, port := range ports {
		if filter == "" {
			filter = fmt.Sprintf("tcp port %d || udp port %d", port, port)
		} else {
			filter = filter + fmt.Sprintf(" or (tcp port %d || udp port %d)", port, port)
		}
	}

	// And apply the filter to our pcap instance
	err := l.handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// Not sure about the layers here
	for _, port := range ports {
		layers.RegisterUDPPortLayerType(layers.UDPPort(port), photon.PhotonLayerType)
		layers.RegisterTCPPortLayerType(layers.TCPPort(port), photon.PhotonLayerType)
	}

	// Overwrite source so we don't have garbage packets
	l.source = gopacket.NewPacketSource(l.handle, l.handle.LinkType())

	// And run the packet parser again.
	l.run()
}

func (l *listener) startOffline(path string) {
	handle, err := pcap.OpenOffline(path)
	if err != err {
		log.Fatalf("Problem creating offline source. Error: %v", err)
	}
	l.handle = handle

	layers.RegisterUDPPortLayerType(5056, photon.PhotonLayerType)
	l.source = gopacket.NewPacketSource(handle, handle.LinkType())

	l.run()
}

func (l *listener) run() {
	l.running = true
	l.stopping = false

	for {
		select {
		case <-l.quit:
			l.running = false

			if l.stopping {
				l.handle.Close()
			}

			return
		case packet := <-l.source.Packets():
			// Why can a packet be nil!?
			if packet != nil {
				l.processPacket(packet)
			}
		}
	}
}

func (l *listener) stop() {
	log.Debugf("Stopping listener (%s)...", l.deviceName)
	l.stopping = true
	if l.running {
		l.quit <- true
	}
}

func (l *listener) processPacket(packet gopacket.Packet) {
	layer := packet.Layer(photon.PhotonLayerType)

	if layer == nil {
		return
	}

	content, _ := layer.(photon.PhotonLayer)

	for _, command := range content.Commands {
		switch command.Type {
		case photon.SendReliableType:
			l.onReliableCommand(&command)
		case photon.SendReliableFragmentType:
			msg, _ := command.ReliableFragment()
			result := l.fragments.Offer(msg)
			if result != nil {
				l.onReliableCommand(result)
			}
		}
	}
}

func (l *listener) onReliableCommand(command *photon.PhotonCommand) {
	msg, _ := command.ReliableMessage()
	params, _ := photon.DecodeReliableMessage(msg)

	switch msg.Type {
	case photon.OperationRequest:
		operation := decodeRequest(params)

		if operation != nil {
			l.router.newOperation <- operation
		}
	case photon.OperationResponse:
		operation := decodeResponse(params)

		if operation != nil {
			l.router.newOperation <- operation
		}
	}
}
