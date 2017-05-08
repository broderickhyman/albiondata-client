package main

import (
	"flag"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/regner/albion-market-data-relay/client/assemblers"
)

func main() {
	log.Print("Starting up AMDR client...")
	deviceName := networkDeviceName()
	handle, err := pcap.OpenLive(*deviceName, 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	var filter = "udp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	source.NoCopy = true

	assembler := assemblers.NewMarketAssembler()

	log.Print("Starting to process packets...")
	for packet := range source.Packets() {
		assembler.ProcessPacket(packet)
	}
}

func networkDeviceName() *string {
	deviceName := flag.String("d", "", "Specifies the network device name. If not specified the first enumerated device will be used.")
	flag.Parse()
	if *deviceName == "" {
		devs, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}
		if len(devs) == 0 {
			log.Fatal("Unable to find network device.")
		}
		*deviceName = devs[0].Name
	}
	return deviceName
}
