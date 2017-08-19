package client

import (
	"github.com/regner/albiondata-client/log"
)

func processOfflinePcap(path string) {
	log.Info("Beginning offline process...")

	r := newRouter()
	go r.run()

	l := newListener(r)
	l.startOffline(path)
}
