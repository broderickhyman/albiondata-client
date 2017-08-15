package client

import (
	"strings"

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
		urls := strings.Split(ConfigGlobal.IngestBaseUrl, ",")
		uploaders := []iuploader{}
		for _, url := range urls {
			if url[0:4] == "http" {
				uploaders = append(uploaders, newHTTPUploader(url))
			} else if url[0:4] == "nats" {
				uploaders = append(uploaders, newNATSUploader(url))
			}
		}

		if len(uploaders) > 1 {
			uploader = newMultiUploader(uploaders)
		} else {
			uploader = uploaders[0]
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
