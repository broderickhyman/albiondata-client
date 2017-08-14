package client

type config struct {
	IngestBaseUrl string
	DisableUpload bool
	SaveLocally   bool
	OfflinePath   string
	Offline       bool
	Debug         bool
	LogLevel      string
	VersionDump   bool
	ListenDevices string
}

var ConfigGlobal = &config{
	LogLevel: "INFO",
}
