package client

import (
	"github.com/regner/albionmarket-client/log"
)

type Router struct {
	albionstate  *albionState
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	return &Router{
		albionstate:  &albionState{},
		newOperation: make(chan operation, 1000),
		quit:         make(chan bool),
	}
}

func (r *Router) run() {
	for {
		select {
		case <-r.quit:
			log.Debug("Closing router...")
			return
		case op := <-r.newOperation:
			op.Process(r.albionstate)
		}
	}
}
