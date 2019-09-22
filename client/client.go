package client

import (
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/broderickhyman/albiondata-client/log"
	colorable "github.com/mattn/go-colorable"
)

var version string

type Client struct {
}

func NewClient(_version string) *Client {
	version = _version
	return &Client{}
}

func (client *Client) Run() {
	if ConfigGlobal.LogToFile {
		log.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true, DisableSorting: true, ForceColors: false})
		f, err := os.OpenFile("albiondata-client-output.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err == nil {
			multiWriter := io.MultiWriter(os.Stdout, f)
			log.SetOutput(multiWriter)
		} else {
			log.SetOutput(os.Stdout)
		}
	} else {
		log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableSorting: true, ForceColors: true})
		log.SetOutput(colorable.NewColorableStdout())
	}

	log.Infof("Starting Albion Data Client, version: %s", version)
	log.Info("This is a third-party application and is in no way affiliated with Sandbox Interactive or Albion Online.")
	log.Info("Additional parameters can listed by calling this file with the -h parameter.")

	createDispatcher()

	if ConfigGlobal.Offline {
		processOffline(ConfigGlobal.OfflinePath)
	} else {
		pw := newProcessWatcher()
		pw.run()
	}
}
