package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/regner/amdr/client/assemblers"
	"log"
)

func main() {
	// Device set to "en0" for local development
	// TODO: Make the device configurable
	handle, err := pcap.OpenLive("eth0", 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	var filter string = "udp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	source.NoCopy = true

	assembler := assemblers.NewMarketAssembler()

	for packet := range source.Packets() {
		assembler.ProcessPacket(packet)
	}
}
