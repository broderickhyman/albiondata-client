package config

type Config struct {
	IngestUrl     string
	DisableUpload bool
	SaveLocally   bool
	OfflinePath   string
	Offline       bool
}

var GlobalConfiguration = &Config{}
