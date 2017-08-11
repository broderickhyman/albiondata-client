package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/regner/albionmarket-client/client/config"
	"github.com/regner/albionmarket-client/log"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Run() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true})
	log.SetLevel(logrus.InfoLevel)

	log.Info("Starting the Albion Market Client...")

	if config.GlobalConfiguration.Offline {
		proccessOfflinePcap(config.GlobalConfiguration.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
