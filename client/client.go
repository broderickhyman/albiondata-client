package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/regner/albiondata-client/log"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Run() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true})

	log.Info("Starting the Albion Data Client...")
	log.Info("This is a third-party application and is in no way affiliated with Sandbox Interactive or Albion Online.")

	createDispatcher()

	if ConfigGlobal.Offline {
		processOfflinePcap(ConfigGlobal.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
