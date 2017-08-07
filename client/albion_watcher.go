package client

import (
	"log"
	"time"

	"github.com/google/gopacket/pcap"
)

type albionProcessWatcher struct {
	pid       int
	known     []int
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
	apw.r.run()

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
	added, removed := diffIntSets(apw.known, current)

	for _, port := range added {
		devices := getDevices()

		for _, device := range devices {
			s := createOnlineSource(device.Name, port)
			l := newListener(s, apw.r)
			go l.run()

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

func getDevices() []pcap.Interface {
	devices, err := pcap.FindAllDevs()

	if err != nil {
		log.Fatal(err)
	}
	if len(devices) == 0 {
		log.Fatal("Unable to find network device.")
	}

	return devices
}
