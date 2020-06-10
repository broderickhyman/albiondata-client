package client

import (
	"runtime"
	"time"

	"github.com/broderickhyman/albiondata-client/log"
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
	log.Debug("Watching processes for Albion to start...")

	for {
		var processName string

		if runtime.GOOS == "windows" {
			processName = "Albion-Online.exe"
		} else {
			processName = "Albion-Online"
		}

		current := findProcess(processName)

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
