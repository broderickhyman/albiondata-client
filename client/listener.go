package client

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	photon "github.com/hmadison/photon_spectator"
)

type listener struct {
	source    *gopacket.PacketSource
	fragments *photon.FragmentBuffer
	quit      chan bool
	router    *Router
}

func newListener(source *gopacket.PacketSource, router *Router) *listener {
	return &listener{
		source:    source,
		fragments: photon.NewFragmentBuffer(),
		quit:      make(chan bool),
		router:    router,
	}
}

func (l *listener) run() {
	log.Printf("Starting listener...")

	for {
		select {
		case <-l.quit:
			log.Printf("Listener shutting down...")

			return
		case packet := <-l.source.Packets():
			// Why can a packet be nil!?
			if packet != nil {
				l.processPacket(packet)
			}
		default:
		}
	}
}

func (l *listener) processPacket(packet gopacket.Packet) {
	log.Print("Processing a packet...")

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
	operation := decode(params)

	if operation != nil {
		l.router.newOperation <- operation
	}
}

func createOnlineSource(device string, port int) *gopacket.PacketSource {
	handle, err := pcap.OpenLive(device, 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	err = handle.SetBPFFilter(fmt.Sprintf("tcp port %d || udp port %d", port, port))
	if err != nil {
		log.Fatal(err)
	}

	layers.RegisterUDPPortLayerType(layers.UDPPort(port), photon.PhotonLayerType)
	layers.RegisterTCPPortLayerType(layers.TCPPort(port), photon.PhotonLayerType)
	return gopacket.NewPacketSource(handle, handle.LinkType())
}
