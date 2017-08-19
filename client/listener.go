package client

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	photon "github.com/hmadison/photon_spectator"
	"github.com/regner/albiondata-client/log"
)

type listener struct {
	handle      *pcap.Handle
	source      *gopacket.PacketSource
	displayName string
	fragments   *photon.FragmentBuffer
	quit        chan bool
	router      *Router
}

func newListener(router *Router) *listener {
	return &listener{
		fragments: photon.NewFragmentBuffer(),
		quit:      make(chan bool, 1),
		router:    router,
	}
}

func (l *listener) startOnline(device string, port int) {
	handle, err := pcap.OpenLive(device, 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	l.handle = handle

	err = l.handle.SetBPFFilter(fmt.Sprintf("tcp port %d || udp port %d", port, port))
	if err != nil {
		log.Fatal(err)
	}

	layers.RegisterUDPPortLayerType(layers.UDPPort(port), photon.PhotonLayerType)
	layers.RegisterTCPPortLayerType(layers.TCPPort(port), photon.PhotonLayerType)
	l.source = gopacket.NewPacketSource(l.handle, l.handle.LinkType())

	l.displayName = fmt.Sprintf("online: %s:%d", device, port)
	l.run()
}

func (l *listener) startOffline(path string) {
	handle, err := pcap.OpenOffline(path)
	if err != err {
		log.Fatalf("Problem creating offline source. Error: %v", err)
	}
	l.handle = handle

	for _, port := range []int{5055, 5056} {
		layers.RegisterUDPPortLayerType(layers.UDPPort(port), photon.PhotonLayerType)
		layers.RegisterTCPPortLayerType(layers.TCPPort(port), photon.PhotonLayerType)
	}
	l.source = gopacket.NewPacketSource(handle, handle.LinkType())

	l.displayName = fmt.Sprintf("offline: %s", path)
	l.run()
}

func (l *listener) run() {
	log.Debugf("Starting listener (%s)...", l.displayName)

	for {
		select {
		case <-l.quit:
			log.Debugf("Listener shutting down (%s)...", l.displayName)
			l.handle.Close()

			return
		case packet := <-l.source.Packets():
			if packet != nil {
				l.processPacket(packet)
			} else {
				// MUST only happen with the offline processor.
				l.handle.Close()
				return
			}
		}
	}
}

func (l *listener) stop() {
	l.quit <- true
	l.handle.Close()
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
	case photon.EventDataType:
	}
}
