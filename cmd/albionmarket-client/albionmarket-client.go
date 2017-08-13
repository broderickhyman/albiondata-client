package main

import (
	"flag"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pcdummy/go-githubupdate/updater"
	"github.com/regner/albionmarket-client/client"
	"github.com/regner/albionmarket-client/client/config"
	"github.com/regner/albionmarket-client/log"
)

var version string

func init() {
	flag.StringVar(
		&config.GlobalConfiguration.IngestBaseUrl,
		"i",
		"https://albion-market.com/api/v1/ingest/",
		"Base URL to send data to.",
	)

	flag.BoolVar(
		&config.GlobalConfiguration.DisableUpload,
		"d",
		false,
		"If specified no attempts will be made to upload data to remote server.",
	)

	flag.BoolVar(
		&config.GlobalConfiguration.SaveLocally,
		"s",
		false,
		"If specified all market orders will be saved locally.",
	)

	flag.StringVar(
		&config.GlobalConfiguration.OfflinePath,
		"o",
		"",
		"Parses a local file instead of checking albion ports.",
	)

	flag.BoolVar(
		&config.GlobalConfiguration.Debug,
		"debug",
		false,
		"Enable debug logging.",
	)

	flag.BoolVar(
		&config.GlobalConfiguration.VersionDump,
		"version",
		false,
		"Print the current version.",
	)
}

func main() {
	flag.Parse()

	if config.GlobalConfiguration.VersionDump {
		log.Infof("albionmarket-client version: %v", version)
		return
	}

	if config.GlobalConfiguration.Debug {
		config.GlobalConfiguration.LogLevel = "DEBUG"
	}

	level, err := logrus.ParseLevel(strings.ToLower(config.GlobalConfiguration.LogLevel))
	if err != nil {
		log.Errorf("Error getting level: %v", err)
	}

	log.SetLevel(level)

	if config.GlobalConfiguration.OfflinePath != "" {
		config.GlobalConfiguration.Offline = true
	}

	// Updater
	if version != "" && !strings.Contains(version, "dev") {
		u := updater.NewUpdater(
			version,
			"regner",
			"albionmarket-client",
			"update-",
		)

		go func() {
			for {
				available, err := u.CheckUpdateAvailable()
				if err != nil {
					log.Errorf("%v", err)
					return
				}

				log.Infof("A new update %s is available", available)
				if available != "" {
					err := u.Update()
					if err != nil {
						log.Errorf("%v", err)
						return
					}

					log.Infof(
						"The update %s has been installed, please restart albionmarket-client.",
						available,
					)
				}

				// Check again in 2 hours
				time.Sleep(time.Hour * 2)
			}
		}()
	}

	c := client.NewClient()
	c.Run()
}
