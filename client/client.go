package client

import (
	"log"

	"github.com/regner/albionmarket-client/client/config"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Run() {
	log.Print("Starting the Albion Market Client...")

	if config.GlobalConfiguration.Offline {
		proccessOfflinePcap(config.GlobalConfiguration.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
