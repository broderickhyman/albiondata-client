package client

import (
	"github.com/Sirupsen/logrus"
	"github.com/broderickhyman/albiondata-client/log"
	"github.com/mattn/go-colorable"
)

var version string

type Client struct {
}

func NewClient(_version string) *Client {
	version = _version
	return &Client{}
}

func (client *Client) Run() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true, ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	log.Infof("Starting Albion Data Client version %s", version)
	log.Info("This is a third-party application and is in no way affiliated with Sandbox Interactive or Albion Online.")

	createDispatcher()

	if ConfigGlobal.Offline {
		processOfflinePcap(ConfigGlobal.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
