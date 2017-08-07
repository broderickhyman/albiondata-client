package client

import (
	"log"
	"time"
)

type processesWatcher struct {
	knownAlbions   []int
	albionWatchers map[int]*albionProcessWatcher
}

func newProcessWatcher() *processesWatcher {
	return &processesWatcher{
		albionWatchers: make(map[int]*albionProcessWatcher),
	}
}

func (pw *processesWatcher) run() {
	log.Print("Watching processes for Albion to start...")

	for {
		current := findProcess("Albion-Online")
		added, removed := diffIntSets(pw.knownAlbions, current)

		for _, pid := range added {
			apw := newAlbionProcessWatcher(pid)

			pw.albionWatchers[pid] = apw

			go apw.run()
		}

		for _, pid := range removed {
			pw.albionWatchers[pid].quit <- true
			delete(pw.albionWatchers, pid)
		}

		pw.knownAlbions = current
		time.Sleep(time.Second)
	}
}
