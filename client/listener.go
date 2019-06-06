package client

import (
	"albiondata-client/log"
	"fmt"

	photon "github.com/broderickhyman/photon_spectator"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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
	params, err := photon.DecodeReliableMessage(msg)
	if err != nil {
		log.Debugf("Error while decoding parameters: %v", err)
		// reset error message
		err = nil
	}

	var operation operation

	switch msg.Type {
	case photon.OperationRequest:
		operation, err = decodeRequest(params)
		if params != nil {
			log.Debugf("OperationRequest: %s", OperationType(params["253"].(int16)))
		}
	case photon.OperationResponse:
		operation, err = decodeResponse(params)
		if params != nil && params["253"] != nil {
			log.Debugf("OperationResponse: %s", OperationType(params["253"].(int16)))
		}
	case photon.EventDataType:
		operation, err = decodeEvent(params)
		if params != nil && params["252"] != nil {
			log.Debugf("EventDataType: %d -- %s", params["252"].(int16), params)
		}
		//  default:
		//    log.Debugf("[%d] (%d) %s (%d)", msg.Type, msg.ParamaterCount, msg.Data, len(msg.Data))
	}

	if err != nil {
		log.Debugf("Error while decoding an event or operation: %v", err)
	}

	if operation != nil {
		l.router.newOperation <- operation
	}
}
