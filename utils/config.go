package utils

import (
	"flag"
	"log"
)

type Config struct {
	IngestUrl     string
	DisableUpload bool
	SaveLocally   bool
	OfflinePath   string
}

func LoadConfig() Config {
	var config = Config{}

	flag.StringVar(&config.IngestUrl, "i",
		"https://albion-market.com/api/v1/ingest/",
		"URL to send market data to.")

	flag.BoolVar(&config.DisableUpload, "d", false,
		"If specified no attempts will be made to upload data to remote server.")

	flag.BoolVar(&config.SaveLocally, "s", false,
		"If specified all market orders will be saved locally.")

	flag.StringVar(&config.OfflinePath, "r", "",
		"Parses a local file instead of checking albion ports")

	flag.Parse()

	if config.OfflinePath == "" {
		if config.DisableUpload {
			 log.Print("Remote upload of market orders is disabled!")
		} else {
				log.Printf("Using the following ingest: %v", config.IngestUrl)
		}
	}

	if config.SaveLocally {
		log.Print("Saving market orders locally.")
	}

	return config
}
