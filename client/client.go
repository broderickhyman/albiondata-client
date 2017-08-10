package client

import (
	"log"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{
	}
}

func (client *Client) Run() {
	log.Print("Starting the Albion Market Client...")

	if GlobalConfiguration.Offline {
		proccessOfflinePcap(GlobalConfiguration.OfflinePath)
	} else {
		pw := newProcessWatcher()
		go pw.run()
	}

	blockForever()
}
