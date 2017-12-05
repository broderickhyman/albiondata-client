package client

import (
	"strings"
	"time"

	"github.com/alexscott/albiondata-client/log"
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
		apw.devices = getAllPhysicalInterface()
		log.Debugf("Will listen to these devices: %v", apw.devices)
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
			l.stop()
		}

		delete(apw.listeners, port)
	}

	apw.r.quit <- true

	log.Printf("Albion watcher closed for PID \"%d\"...", apw.pid)
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
