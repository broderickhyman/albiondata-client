package client

import (
	"strings"
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/regner/albionmarket-client/log"
)

type albionProcessWatcher struct {
	pid       int
	known     []int
	devices   []string
	listeners map[int][]*listener
	quit      chan bool
	r         *Router
}

func newAlbionProcessWatcher(pid int) *albionProcessWatcher {
	return &albionProcessWatcher{
		pid:       pid,
		listeners: make(map[int][]*listener),
		quit:      make(chan bool),
		r:         newRouter(),
	}
}

func (apw *albionProcessWatcher) run() {
	log.Printf("Watching Albion process with PID \"%d\"...", apw.pid)
	if ConfigGlobal.ListenDevices != "" {
		apw.devices = strings.Split(ConfigGlobal.ListenDevices, ",")
	} else {
		apw.devices = apw.getDevices()
	}
	go apw.r.run()

	for {
		select {
		case <-apw.quit:
			apw.closeWatcher()
			return
		default:
			apw.updateListeners()
			time.Sleep(time.Second)
		}
	}
}

func (apw *albionProcessWatcher) closeWatcher() {
	log.Printf("Albion watcher closed for PID \"%d\"...", apw.pid)

	for port := range apw.listeners {
		for _, l := range apw.listeners[port] {
			l.quit <- true
		}

		delete(apw.listeners, port)
	}

	apw.r.quit <- true
}

func (apw *albionProcessWatcher) updateListeners() {
	current := getProcessPorts(apw.pid)
	filtered := []int{}
	for _, port := range current {
		if port == 0 {
			continue
		}

		filtered = append(filtered, port)
	}

	added, removed := diffIntSets(apw.known, filtered)

	for _, port := range added {
		for _, device := range apw.devices {
			l := newListener(apw.r)
			go l.startOnline(device, port)

			apw.listeners[port] = append(apw.listeners[port], l)
		}
	}

	for _, port := range removed {
		for _, l := range apw.listeners[port] {
			l.quit <- true
		}

		delete(apw.listeners, port)
	}

	apw.known = current
}

func (apw *albionProcessWatcher) getDevices() []string {
	devices, err := pcap.FindAllDevs()

	// Filter out devices that we aren't able to listen to.
	// they bring error's like "NFLOG link-layer type filtering not implemented"
	blacklisted := []string{"nflog", "nfqueue", "usbmon1", "usbmon2", "usbmon3", "usbmon4", "oracle"}

	for _, bl := range blacklisted {
		found := false
		id1 := -1
		for id2, device := range devices {
			if device.Name == bl {
				found = true
				id1 = id2
				break
			}
		}
		if found {
			devices = append(devices[:id1], devices[id1+1:]...)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	if len(devices) == 0 {
		log.Fatal("Unable to find network device.")
	}

	strDevices := []string{}
	for _, dev := range devices {
		strDevices = append(strDevices, dev.Name)
	}

	return strDevices
}
