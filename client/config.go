package client

type config struct {
	AllowedWSHosts        []string
	Debug                 bool
	DisableUpload         bool
	EnableWebsockets      bool
	ListenDevices         string
	LogLevel              string
	LogToFile             bool
	Minimize              bool
	Offline               bool
	OfflinePath           string
	RecordPath            string
	PrivateIngestBaseUrls string
	PublicIngestBaseUrls  string
}

var ConfigGlobal = &config{
	LogLevel: "INFO",
}
