package client

import (
	"github.com/broderickhyman/albiondata-client/log"
)

type Router struct {
	albionstate  *albionState
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	return &Router{
		albionstate:  &albionState{LocationId: -1},
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
			go op.Process(r.albionstate)
		}
	}
}
