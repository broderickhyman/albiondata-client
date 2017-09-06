package client

import (
	"os"

	"github.com/regner/albiondata-client/log"
)

func processOfflinePcap(path string) {
	log.Info("Beginning offline process...")

	r := newRouter()
	go r.run()

	_, err := os.Stat(path)

	if err != nil {
		log.Error("Could not find {}: ", path, err)

		return
	}

	l := newListener(r)
	l.startOffline(path)
}
