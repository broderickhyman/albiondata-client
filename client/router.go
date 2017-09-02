package client

import (
	"strings"

	"github.com/regner/albiondata-client/log"
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
		// PUBLIC URLs
		urls := strings.Split(ConfigGlobal.PublicIngestBaseUrls, ",")
		uploaders := []iuploader{}
		for _, url := range urls {
			if len(url) < 4 {
				continue
			}

			if url[0:4] == "http" {
				uploaders = append(uploaders, newHTTPUploader(url, false))
			} else if url[0:4] == "nats" {
				uploaders = append(uploaders, newNATSUploader(url, false))
			} else if url[0:4] == "noop" {
				uploaders = append(uploaders, newNOOPUploader(false))
			}
		}

		// Private URLs
		urls = strings.Split(ConfigGlobal.PrivateIngestBaseUrls, ",")
		for _, url := range urls {
			if len(url) < 4 {
				continue
			}

			if url[0:4] == "http" {
				uploaders = append(uploaders, newHTTPUploader(url, true))
			} else if url[0:4] == "nats" {
				uploaders = append(uploaders, newNATSUploader(url, true))
			} else if url[0:4] == "noop" {
				uploaders = append(uploaders, newNOOPUploader(true))
			}
		}

		uploader = newMultiUploader(uploaders)
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
