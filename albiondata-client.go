package main

import (
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/broderickhyman/albiondata-client/client"
	"github.com/broderickhyman/albiondata-client/log"
	"github.com/broderickhyman/albiondata-client/systray"
	"github.com/broderickhyman/go-githubupdate/updater"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var version string

func init() {
	// Setup the config file and parse values
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	// if we cannot find the configuration file, set Websockets to false
	if err != nil {
		viper.Set("EnableWebsockets", false)
	}

	client.ConfigGlobal.EnableWebsockets = viper.GetBool("EnableWebsockets")
	client.ConfigGlobal.AllowedWSHosts = viper.GetStringSlice("AllowedWebsocketHosts")

	flag.BoolVar(
		&client.ConfigGlobal.Debug,
		"debug",
		false,
		"Enable debug logging.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.DisableUpload,
		"d",
		false,
		"If specified no attempts will be made to upload data to remote server.",
	)

	flag.StringVar(
		&client.ConfigGlobal.ListenDevices,
		"l",
		"",
		"Listen on this comma separated devices instead of all available",
	)

	flag.BoolVar(
		&client.ConfigGlobal.LogToFile,
		"output-file",
		false,
		"Enable logging to file.",
	)

	flag.StringVar(
		&client.ConfigGlobal.OfflinePath,
		"o",
		"",
		"Parses a local file instead of checking albion ports.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.Minimize,
		"minimize",
		false,
		"Automatically minimize the window.",
	)

	flag.StringVar(
		&client.ConfigGlobal.PublicIngestBaseUrls,
		"i",
		"http+pow://www.albion-online-data.com:4223",
		"Base URL to send PUBLIC data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders. Comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.PrivateIngestBaseUrls,
		"p",
		"",
		"Base URL to send PRIVATE data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders. Comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.RecordPath,
		"record",
		"",
		"Enable recording commands to a file for debugging later.",
	)

	flag.StringVar(
		&client.ConfigGlobal.DebugEventsString,
		"events",
		"",
		"Whitelist of event IDs to output messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.DebugEventsBlacklistString,
		"events-ignore",
		"",
		"Blacklist of event IDs to hide messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.DebugOperationsString,
		"operations",
		"",
		"Whitelist of operation IDs to output messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&client.ConfigGlobal.DebugOperationsBlacklistString,
		"operations-ignore",
		"",
		"Blacklist of operation IDs to hide messages when debugging. Comma separated.",
	)

	flag.BoolVar(
		&client.ConfigGlobal.DebugIgnoreDecodingErrors,
		"ignore-decode-errors",
		false,
		"Ignore the decoding errors when debugging",
	)
}

func main() {
	flag.Parse()

	if client.ConfigGlobal.Debug {
		client.ConfigGlobal.LogLevel = "DEBUG"
	}
	client.ConfigGlobal.DebugEvents = make(map[int]bool)
	if client.ConfigGlobal.DebugEventsString != "" {
		for _, event := range strings.Split(client.ConfigGlobal.DebugEventsString, ",") {
			number, err := strconv.Atoi(event)
			if err == nil {
				client.ConfigGlobal.DebugEvents[number] = true
			}
		}
	}
	if client.ConfigGlobal.DebugEventsBlacklistString != "" {
		for _, event := range strings.Split(client.ConfigGlobal.DebugEventsBlacklistString, ",") {
			number, err := strconv.Atoi(event)
			if err == nil {
				client.ConfigGlobal.DebugEvents[number] = false
			}
		}
	}

	client.ConfigGlobal.DebugOperations = make(map[int]bool)
	if client.ConfigGlobal.DebugOperationsString != "" {
		for _, operation := range strings.Split(client.ConfigGlobal.DebugOperationsString, ",") {
			number, err := strconv.Atoi(operation)
			if err == nil {
				client.ConfigGlobal.DebugOperations[number] = true
			}
		}
	}

	if client.ConfigGlobal.DebugOperationsBlacklistString != "" {
		for _, operation := range strings.Split(client.ConfigGlobal.DebugOperationsBlacklistString, ",") {
			number, err := strconv.Atoi(operation)
			if err == nil {
				client.ConfigGlobal.DebugOperations[number] = false
			}
		}
	}

	level, err := logrus.ParseLevel(strings.ToLower(client.ConfigGlobal.LogLevel))
	if err != nil {
		log.Errorf("Error getting level: %v", err)
	}

	log.SetLevel(level)

	if client.ConfigGlobal.OfflinePath != "" {
		client.ConfigGlobal.Offline = true
		client.ConfigGlobal.DisableUpload = true
	}

	if client.ConfigGlobal.DisableUpload {
		log.Info("Upload is disabled.")
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
			maxTries := 2
			for i := 0; i < maxTries; i++ {
				err := u.BackgroundUpdater()
				if err != nil {
					log.Error(err.Error())
					log.Info("Will try again in 60 seconds. You may need to run the client as Administrator.")
					// Sleep and hope the network connects
					time.Sleep(time.Second * 60)
				} else {
					break
				}
			}
		}()
	}
}
