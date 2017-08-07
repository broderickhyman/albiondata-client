package main

import (
	"flag"

	"github.com/regner/albionmarket-client/client"
)

var (
	options = &client.Config{}
)

func init() {
	flag.StringVar(
		&options.IngestUrl,
		"i",
		"https://albion-market.com/api/v1/ingest/",
		"URL to send market data to.",
	)

	flag.BoolVar(
		&options.DisableUpload,
		"d",
		false,
		"If specified no attempts will be made to upload data to remote server.",
	)

	flag.BoolVar(
		&options.SaveLocally,
		"s",
		false,
		"If specified all market orders will be saved locally.",
	)

	flag.StringVar(
		&options.OfflinePath,
		"o",
		"",
		"Parses a local file instead of checking albion ports",
	)
}

func main() {
	flag.Parse()

	if options.OfflinePath != "" {
		options.Offline = true
	}

	c := client.NewClient(options)
	c.Run()
}
