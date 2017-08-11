package main

import (
	"flag"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/regner/albionmarket-client/client"
	"github.com/regner/albionmarket-client/client/config"
	"github.com/regner/albionmarket-client/log"
)

func init() {
	flag.StringVar(
		&config.GlobalConfiguration.IngestUrl,
		"i",
		"https://albion-market.com/api/v1/ingest/",
		"URL to send market data to.",
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
		"Parses a local file instead of checking albion ports",
	)

	flag.BoolVar(
		&config.GlobalConfiguration.Debug,
		"debug",
		false,
		"Enable debug logging",
	)
}

func main() {
	flag.Parse()

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

	c := client.NewClient()
	c.Run()
}
