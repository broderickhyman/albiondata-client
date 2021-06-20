package client

import (
	"flag"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/broderickhyman/albiondata-client/log"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	AllowedWSHosts                 []string
	Debug                          bool
	DebugEvents                    map[int]bool
	DebugEventsString              string
	DebugEventsBlacklistString     string
	DebugOperations                map[int]bool
	DebugOperationsString          string
	DebugOperationsBlacklistString string
	DebugIgnoreDecodingErrors      bool
	DisableUpload                  bool
	EnableWebsockets               bool
	ListenDevices                  string
	LogLevel                       string
	LogToFile                      bool
	Minimize                       bool
	Offline                        bool
	OfflinePath                    string
	RecordPath                     string
	PrivateIngestBaseUrls          string
	PublicIngestBaseUrls           string
	NoCPULimit                     bool
}

//config global config data
var ConfigGlobal = &config{
	LogLevel: "INFO",
}

func (config *config) Setup() {
	config.setupWebsocketFlags()
	config.setupDebugFlags()
	config.setupCommonFlags()

	flag.Parse()

	if config.OfflinePath != "" {
		config.Offline = true
		config.DisableUpload = true
	}

	if config.DisableUpload {
		log.Info("Upload is disabled.")
	}
	config.setupLogs()
	config.setupDebugEvents()
	config.setupDebugOperations()
}

func (config *config) setupWebsocketFlags() {
	// Setup the config file and parse values
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	// if we cannot find the configuration file, set Websockets to false
	if err != nil {
		viper.Set("EnableWebsockets", false)
	}

	config.EnableWebsockets = viper.GetBool("EnableWebsockets")
	config.AllowedWSHosts = viper.GetStringSlice("AllowedWebsocketHosts")
}

func (config *config) setupDebugFlags() {
	flag.BoolVar(
		&config.Debug,
		"debug",
		false,
		"Enable debug logging.",
	)

	flag.StringVar(
		&config.DebugEventsString,
		"events",
		"",
		"Whitelist of event IDs to output messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&config.DebugEventsBlacklistString,
		"events-ignore",
		"",
		"Blacklist of event IDs to hide messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&config.DebugOperationsString,
		"operations",
		"",
		"Whitelist of operation IDs to output messages when debugging. Comma separated.",
	)

	flag.StringVar(
		&config.DebugOperationsBlacklistString,
		"operations-ignore",
		"",
		"Blacklist of operation IDs to hide messages when debugging. Comma separated.",
	)

	flag.BoolVar(
		&config.DebugIgnoreDecodingErrors,
		"ignore-decode-errors",
		false,
		"Ignore the decoding errors when debugging",
	)

	flag.BoolVar(
		&config.NoCPULimit,
		"no-limit",
		false,
		"Use all available CPU cores",
	)

}

func (config *config) setupCommonFlags() {
	flag.BoolVar(
		&config.DisableUpload,
		"d",
		false,
		"If specified no attempts will be made to upload data to remote server.",
	)

	flag.StringVar(
		&config.ListenDevices,
		"l",
		"",
		"Listen on this comma separated devices instead of all available",
	)

	flag.BoolVar(
		&config.LogToFile,
		"output-file",
		false,
		"Enable logging to file.",
	)

	flag.StringVar(
		&config.OfflinePath,
		"o",
		"",
		"Parses a local file instead of checking albion ports.",
	)

	flag.BoolVar(
		&config.Minimize,
		"minimize",
		false,
		"Automatically minimize the window.",
	)

	flag.StringVar(
		&config.PublicIngestBaseUrls,
		"i",
		"http+pow://www.albion-online-data.com:4223",
		"Base URL to send PUBLIC data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders. Comma separated.",
	)

	flag.StringVar(
		&config.PrivateIngestBaseUrls,
		"p",
		"",
		"Base URL to send PRIVATE data to, can be 'nats://', 'http://' or 'noop' and can have multiple uploaders. Comma separated.",
	)

	flag.StringVar(
		&config.RecordPath,
		"record",
		"",
		"Enable recording commands to a file for debugging later.",
	)
}

func (config *config) setupLogs() {
	if config.Debug {
		config.LogLevel = "DEBUG"
	}

	level, err := logrus.ParseLevel(strings.ToLower(config.LogLevel))
	if err != nil {
		log.Errorf("Error getting level: %v", err)
	}

	log.SetLevel(level)

	if config.LogToFile {
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
}

func (config *config) setupDebugEvents() {
	config.DebugEvents = make(map[int]bool)
	if config.DebugEventsString != "" {
		for _, event := range strings.Split(config.DebugEventsString, ",") {
			number, err := strconv.Atoi(event)
			if err == nil {
				config.DebugEvents[number] = true
			}
		}
	}
	if config.DebugEventsBlacklistString != "" {
		for _, event := range strings.Split(config.DebugEventsBlacklistString, ",") {
			number, err := strconv.Atoi(event)
			if err == nil {
				config.DebugEvents[number] = false
			}
		}
	}

	// Looping through map keys is purposefully random by design in Go
	for number, shouldDebug := range config.DebugEvents {
		verb := "Ignoring"
		if shouldDebug {
			verb = "Showing"
		}
		log.Debugf("[%v] event: [%v]%v", verb, number, EventType(number))
	}

}

func (config *config) setupDebugOperations() {
	config.DebugOperations = make(map[int]bool)
	if config.DebugOperationsString != "" {
		for _, operation := range strings.Split(config.DebugOperationsString, ",") {
			number, err := strconv.Atoi(operation)
			if err == nil {
				config.DebugOperations[number] = true
			}
		}
	}

	if config.DebugOperationsBlacklistString != "" {
		for _, operation := range strings.Split(config.DebugOperationsBlacklistString, ",") {
			number, err := strconv.Atoi(operation)
			if err == nil {
				config.DebugOperations[number] = false
			}
		}
	}

	// Looping through map keys is purposefully random by design in Go
	for number, shouldDebug := range config.DebugOperations {
		verb := "Ignoring"
		if shouldDebug {
			verb = "Showing"
		}
		log.Debugf("[%v] operation: [%v]%v", verb, number, OperationType(number))
	}

}
