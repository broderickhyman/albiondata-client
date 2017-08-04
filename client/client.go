package client

import (
	"log"
)

type Client struct {
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
	}
}

func (client *Client) Run() {
	log.Print("Starting the Albion Market Client...")

	if client.config.Offline {
		proccessOfflinePcap(client.config.OfflinePath)
	} else {
		pw := newProcessWatcher()
		go pw.run()
	}

	blockForever()
}
