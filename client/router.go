package client

import "log"

type Router struct {
	locationId   int
	newOperation chan operation
	quit         chan bool
}

func newRouter() *Router {
	return &Router{
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
			op.Process(r)
		}
	}
}
