package client

import (
	"github.com/regner/albionmarket-client/log"
)

type Router struct {
	albionstate  *albionState
	uploader     iuploader
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	uploader := iuploader(&noopUploader{})
	if !ConfigGlobal.DisableUpload {
		if ConfigGlobal.IngestBaseUrl[0:4] == "http" {
			uploader = newHTTPUploader(ConfigGlobal.IngestBaseUrl)
		} else if ConfigGlobal.IngestBaseUrl[0:4] == "nats" {
			uploader = newNATSUploader(ConfigGlobal.IngestBaseUrl)
		}
	}

	return &Router{
		albionstate:  &albionState{},
		uploader:     uploader,
		newOperation: make(chan operation, 1000),
		quit:         make(chan bool, 1),
	}
}

func (r *Router) run() {
	for {
		select {
		case <-r.quit:
			log.Debug("Closing router...")
			return
		case op := <-r.newOperation:
			go op.Process(r.albionstate, r.uploader)
		}
	}
}
