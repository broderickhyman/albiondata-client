package client

import (
	"io"
	"os"

	"github.com/broderickhyman/albiondata-client/log"
	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var version string

//Client struct base
type Client struct {
}

//NewClient return a new Client instance
func NewClient(_version string) *Client {
	version = _version
	return &Client{}
}

//Run starts client settings and run
func (client *Client) Run() error {
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

	// Looping through map keys is purposefully random by design in Go
	for number, shouldDebug := range ConfigGlobal.DebugEvents {
		verb := "Ignoring"
		if shouldDebug {
			verb = "Showing"
		}
		log.Debugf("[%v] event: [%v]%v", verb, number, EventType(number))
	}

	// Looping through map keys is purposefully random by design in Go
	for number, shouldDebug := range ConfigGlobal.DebugOperations {
		verb := "Ignoring"
		if shouldDebug {
			verb = "Showing"
		}
		log.Debugf("[%v] operation: [%v]%v", verb, number, OperationType(number))
	}

	createDispatcher()

	if ConfigGlobal.Offline {
		processOffline(ConfigGlobal.OfflinePath)
	} else {
		apw := newAlbionProcessWatcher()
		return apw.run()
	}
	return nil
}
