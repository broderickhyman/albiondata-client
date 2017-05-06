package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"github.com/regner/amdr-client/marketassembley"
)

func main() {
	log.Print("Starting things...")
	// Device set to "en0" for local development
	// TODO: Make the device configurable
	handle, err := pcap.OpenLive("en0", 2048, false, pcap.BlockForever)
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

	assembler := marketassembley.NewMarketAssembler()

	log.Println("Starting to read packets...")
	for packet := range source.Packets() {
		assembler.ProcessPacket(packet)
	}
}