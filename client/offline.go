package client

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	photon "github.com/hmadison/photon_spectator"
)

func proccessOfflinePcap(path string) {
	log.Print("Beginning offline process...")

	r := newRouter()
	go r.run()

	s := createOfflineSource(path)

	l := newListener(s, r)
	l.run()
}

func createOfflineSource(path string) *gopacket.PacketSource {
	handle, err := pcap.OpenOffline(path)
	if err != err {
		log.Fatalf("Problem creating offline source. Error: %v", err)
	}

	layers.RegisterUDPPortLayerType(5056, photon.PhotonLayerType)
	return gopacket.NewPacketSource(handle, handle.LinkType())
}
