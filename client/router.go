package client

import (
	"encoding/gob"
	"os"

	"github.com/broderickhyman/albiondata-client/log"
	photon "github.com/broderickhyman/photon_spectator"
)

//Router struct definitions
type Router struct {
	albionstate         *albionState
	newOperation        chan operation
	recordPhotonCommand chan photon.PhotonCommand
	quit                chan bool
}

func newRouter() *Router {
	return &Router{
		albionstate:         &albionState{LocationId: -1},
		newOperation:        make(chan operation, 1000),
		recordPhotonCommand: make(chan photon.PhotonCommand, 1000),
		quit:                make(chan bool, 1),
	}
}

func (r *Router) run() {
	var encoder *gob.Encoder
	var file *os.File
	if ConfigGlobal.RecordPath != "" {
		file, err := os.Create(ConfigGlobal.RecordPath)
		if err != nil {
			log.Errorf("Could not open commands output file ", err)
		} else {
			encoder = gob.NewEncoder(file)
		}
	}

	for {
		select {
		case <-r.quit:
			log.Debug("Closing router...")
			if file != nil {
				err := file.Close()
				if err != nil {
					log.Errorf("Could not close commands output file ", err)
				}
			}
			return
		case op := <-r.newOperation:
			go op.Process(r.albionstate)
		case command := <-r.recordPhotonCommand:
			if encoder != nil {
				err := encoder.Encode(command)
				if err != nil {
					log.Errorf("Could not encode command ", err)
				}
			}
		}
	}
}
