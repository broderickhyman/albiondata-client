package client

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
}

var ConfigGlobal = &config{
	LogLevel: "INFO",
}
