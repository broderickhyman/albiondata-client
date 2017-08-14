package client

import (
	"time"

	"github.com/google/gopacket/pcap"
	"github.com/regner/albionmarket-client/log"
)

type albionProcessWatcher struct {
	pid       int
	listeners map[string]*listener
	quit      chan bool
	r         *Router
	devices   []string
	lastPorts []int
}

func newAlbionProcessWatcher(pid int) *albionProcessWatcher {
	return &albionProcessWatcher{
		pid:       pid,
		listeners: make(map[string]*listener),
		quit:      make(chan bool),
		r:         newRouter(),
		lastPorts: []int{},
	}
}

func (apw *albionProcessWatcher) run() {
	log.Printf("Watching Albion process with PID \"%d\"...", apw.pid)
	apw.devices = apw.getDevices()
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

	for _, l := range apw.listeners {
		l.stop()
	}

	for _, device := range apw.devices {
		delete(apw.listeners, device)
	}

	apw.r.quit <- true
}

func (apw *albionProcessWatcher) updateListeners() {
	ports := getProcessPorts(apw.pid)

	myPorts := []int{}
	for _, p := range ports {
		if p > 0 {
			myPorts = append(myPorts, p)
		}
	}

	// No need to change ports if they didn't change
	added, removed := diffIntSets(myPorts, apw.lastPorts)
	if len(added) == 0 && len(removed) == 0 {
		return
	}
	apw.lastPorts = myPorts

	for _, device := range apw.devices {
		if _, ok := apw.listeners[device]; !ok {
			l := newListener(apw.r)
			l.createOnline(device)

			apw.listeners[device] = l
		}

		l, _ := apw.listeners[device]

		if len(myPorts) > 0 {
			go l.setPortsAndRun(device, myPorts)
		} else {
			l.stop()
		}

	}
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

	s := []string{}
	for _, dev := range devices {
		s = append(s, dev.Name)
	}

	return s
}
