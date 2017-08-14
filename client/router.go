package client

import (
	"github.com/regner/albionmarket-client/log"
)

type Router struct {
	albionstate  *albionState
	uploader     *uploader
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	return &Router{
		albionstate:  &albionState{},
		uploader:     newUploader(),
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
