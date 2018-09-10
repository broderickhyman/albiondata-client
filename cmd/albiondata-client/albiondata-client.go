package main

import (
	"flag"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/broderickhyman/albiondata-client/client"
	"github.com/broderickhyman/albiondata-client/log"
	"github.com/broderickhyman/albiondata-client/systray"
	"github.com/broderickhyman/go-githubupdate/updater"
)

var version string

func init() {
	flag.StringVar(
		&client.ConfigGlobal.PublicIngestBaseUrls,
		"i",
		"nats://public:thenewalbiondata@www.albion-online-data.com:4222",
		"Base URL to send PUBLIC data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.PrivateIngestBaseUrls,
		"p",
		"",
		"Base URL to send PRIVATE data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders comma separated.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.DisableUpload,
		"d",
		false,
		"If specified no attempts will be made to upload data to remote server.",
	)

	flag.StringVar(
		&client.ConfigGlobal.OfflinePath,
		"o",
		"",
		"Parses a local file instead of checking albion ports.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.Debug,
		"debug",
		false,
		"Enable debug logging.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.LogToFile,
		"output-file",
		false,
		"Enable logging to file.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.VersionDump,
		"version",
		false,
		"Print the current version.",
	)

	flag.StringVar(
		&client.ConfigGlobal.ListenDevices,
		"l",
		"",
		"Listen on this comma separated devices instead of all available",
	)
}

func main() {
	flag.Parse()

	if client.ConfigGlobal.VersionDump {
		log.Infof("albiondata-client version: %v", version)
		return
	}

	if client.ConfigGlobal.Debug {
		client.ConfigGlobal.LogLevel = "DEBUG"
	}

	level, err := logrus.ParseLevel(strings.ToLower(client.ConfigGlobal.LogLevel))
	if err != nil {
		log.Errorf("Error getting level: %v", err)
	}

	log.SetLevel(level)

	if client.ConfigGlobal.OfflinePath != "" {
		client.ConfigGlobal.Offline = true
	}

	startUpdater()

	go systray.Run()

	c := client.NewClient(version)
	c.Run()
}

func startUpdater() {
	if version != "" && !strings.Contains(version, "dev") {
		u := updater.NewUpdater(
			version,
			"broderickhyman",
			"albiondata-client",
			"update-",
		)

		go func() {
			for {
				maxTries := 6
				for i := 0; i < maxTries; i++ {
					err := u.BackgroundUpdater()
					if err != nil {
						if i == maxTries-1 {
							log.Error(err.Error())
						} else {
							// Sleep and hope the network connects
							time.Sleep(time.Second * 10)
						}
					} else {
						break
					}
				}

				// Check again in 2 hours
				time.Sleep(time.Hour * 2)
			}
		}()
	}
}
