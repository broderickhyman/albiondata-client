package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/regner/albionmarket-client/log"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Run() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true})

	log.Info("Starting the Albion Market Client...")

	if ConfigGlobal.Offline {
		processOfflinePcap(ConfigGlobal.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
