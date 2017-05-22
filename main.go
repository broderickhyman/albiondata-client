package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/regner/amdr-client/assemblers"
	"github.com/regner/amdr-client/utils"
)

func main() {
	log.Print("Starting up AMDR client...")
	config := utils.ClientConfig{}

	flag.StringVar(&config.DeviceName, "d", "", "Specifies the network device name. If not specified the first enumerated device will be used.")
	flag.StringVar(&config.IngestUrl, "i", "https://albion-market.com/api/v1/ingest/", "URL to send market data to.")
	flag.Parse()

	config.DeviceName = networkDeviceName(config.DeviceName)

	log.Printf("Using the following network device: %v", config.DeviceName)

	handle, err := pcap.OpenLive(config.DeviceName, 2048, false, pcap.BlockForever)
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

	assembler := assemblers.NewMarketAssembler(config)

	log.Print("Starting to process packets...")
	for packet := range source.Packets() {
		assembler.ProcessPacket(packet)
	}
}

func networkDeviceName(deviceName string) string {
	if deviceName == "" {
		devs, err := pcap.FindAllDevs()
		if err != nil {
			log.Fatal(err)
		}
		if len(devs) == 0 {
			log.Fatal("Unable to find network device.")
		}

		if runtime.GOOS == "windows" {
			for _, device := range devs {
				// Quick and dirt hack around dealing with VirtualBox interfaces on windows
				// as one of them is often the first in the device list
				if device.Description != "Oracle"{
					return device.Name
				}
			}
		}

		return devs[0].Name
	}

	return deviceName
}
