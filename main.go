package main

import (
	"fmt"
	"log"
	"time"

	"github.com/regner/albionmarket-client/asserts"
	"github.com/regner/albionmarket-client/operations"
	"github.com/regner/albionmarket-client/utils"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	photon "github.com/hmadison/photon_spectator"
)

var onOperationSend chan interface{} = make(chan interface{})

func main() {
	log.Print("Starting the Albion Market Client...")
	var conf = utils.LoadConfig()

	if conf.OfflinePath != "" {
		log.Print("Parsing offline pcap file ", conf.OfflinePath)
		go listenToSource(createOfflineSource(conf.OfflinePath), make(chan bool))
		go processOperations()
		time.Sleep(time.Second)
		return
	}

	go processOperations()
	go watchProcesses()

	for {
		time.Sleep(time.Millisecond * 500)
	}
}

func watchProcesses() {
	log.Print("Watching processes for Albion to start")

	var known []int
	var workers = map[int]chan bool{}

	for {
		var current []int = utils.FindProcess("Albion-Online")
		added, removed := utils.DiffIntSets(known, current)

		for _, pid := range added {
			workers[pid] = make(chan bool)
			go watchAlbion(pid, workers[pid])
		}

		for _, pid := range removed {
			log.Printf("Albion process closed (%d)", pid)
			workers[pid] <- true
		}

		known = current
		time.Sleep(time.Second * 5)
	}
}

func watchAlbion(pid int, quit chan bool) {
	log.Printf("Watching Albion (%d)", pid)

	var known []int
	var workers = make(map[int]chan bool)

	for !utils.IsChannelSignaled(quit) {
		var current []int = utils.GetProcessPorts(pid)
		added, removed := utils.DiffIntSets(known, current)

		for _, port := range added {
			log.Printf("Watching Albion (%d) for packets on port %d", pid, port)
			workers[port] = make(chan bool)
			go listenToSource(createOnlineSource(port), workers[port])
		}

		for _, port := range removed {
			log.Printf("Albion (%d) closed port %d", pid, port)
			workers[port] <- true
		}

		known = current
		time.Sleep(time.Second)
	}

	log.Print("Albion watcher closed ", pid)
}

func createOnlineSource(port int) *gopacket.PacketSource {
	handle, err := pcap.OpenLive("en0", 2048, false, pcap.BlockForever)
	asserts.NoError(err)

	err = handle.SetBPFFilter(fmt.Sprintf("tcp port %d || udp port %d", port, port))
	asserts.NoError(err)

	layers.RegisterUDPPortLayerType(layers.UDPPort(port), photon.PhotonLayerType)
	layers.RegisterTCPPortLayerType(layers.TCPPort(port), photon.PhotonLayerType)
	return gopacket.NewPacketSource(handle, handle.LinkType())
}

func createOfflineSource(path string) *gopacket.PacketSource {
	handle, err := pcap.OpenOffline(path)
	asserts.NoError(err)

	layers.RegisterUDPPortLayerType(5056, photon.PhotonLayerType)
	return gopacket.NewPacketSource(handle, handle.LinkType())
}

func listenToSource(source *gopacket.PacketSource, quit chan bool) {
	var packetChannel = source.Packets()
	var fragments = photon.NewFragmentBuffer()

	for !utils.IsChannelSignaled(quit) {
		packet := utils.PollChannel(packetChannel)

		if packet == nil {
			continue
		}

		layer := packet.Layer(photon.PhotonLayerType)

		if layer == nil {
			return
		}

		content, _ := layer.(photon.PhotonLayer)

		for _, command := range content.Commands {
			switch command.Type {
			case photon.SendReliableType:
				onReliableCommand(&command)

			case photon.SendReliableFragmentType:
				msg, _ := command.ReliableFragment()
				result := fragments.Offer(msg)
				if result != nil {
					onReliableCommand(result)
				}
			}
		}
	}
}

func onReliableCommand(command *photon.PhotonCommand) {
	msg, _ := command.ReliableMessage()
	params, _ := photon.DecodeReliableMessage(msg)
	operation := operations.Decode(params)

	if operation != nil {
		onOperationSend <- operation
	}
}

func processOperations() {
	for {
		v := <-onOperationSend
		log.Print("Operation received", v)
	}
}
