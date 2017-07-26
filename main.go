package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/regner/albionmarket-client/assemblers"
	"github.com/regner/albionmarket-client/utils"
)

func main() {
	log.Print("Starting the Albion Market Client...")
	config := utils.ClientConfig{}

	flag.StringVar(&config.DeviceName, "d", "", "Specifies the network device name. If not specified the first enumerated device will be used.")
	flag.StringVar(&config.IngestUrl, "i", "https://albion-market.com/api/v1/ingest/", "URL to send market data to.")
	flag.BoolVar(&config.DisableUpload, "u", false, "If specified no attempts will be made to upload data to remote server.")
	flag.BoolVar(&config.SaveLocally, "s", false, "If specified all market orders will be saved locally.")
	flag.Parse()

	config.DeviceName = networkDeviceName(config.DeviceName)

	log.Printf("Using the following network device: %v", config.DeviceName)

	if config.DisableUpload {
		log.Print("Remote upload of market orders is disabled!")
	} else {
		log.Printf("Using the following ingest: %v", config.IngestUrl)
	}

	if config.SaveLocally {
		log.Print("Saving market orders locally.")
	}

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
