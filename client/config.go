package client

type config struct {
	PublicIngestBaseUrls  string
	PrivateIngestBaseUrls string
	DisableUpload         bool
	OfflinePath           string
	RecordPath            string
	Offline               bool
	Debug                 bool
	LogToFile             bool
	LogLevel              string
	VersionDump           bool
	ListenDevices         string
	EnableWebsockets      bool
	AllowedWSHosts        []string
}

var ConfigGlobal = &config{
	LogLevel: "INFO",
}
