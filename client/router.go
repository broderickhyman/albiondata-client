package client

import (
	"log"

	"github.com/regner/albionmarket-client/client/albionstate"
)

type Router struct {
	state        *albionstate.AlbionState
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	return &Router{
		state:        &albionstate.AlbionState{},
		newOperation: make(chan operation, 1000),
		quit:         make(chan bool),
	}
}

func (r *Router) run() {
	for {
		select {
		case <-r.quit:
			log.Print("Closing router...")
			return
		case op := <-r.newOperation:
			op.Process(r.state)
		}
	}
}
