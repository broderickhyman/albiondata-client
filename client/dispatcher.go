package client

import (
	"encoding/json"
	"net/http"

	"strings"

	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type dispatcher struct {
	publicUploaders  []uploader
	privateUploaders []uploader
}

var (
	wsHub *WSHub
	dis   *dispatcher
)

func createDispatcher() {
	dis = &dispatcher{
		publicUploaders:  createUploaders(strings.Split(ConfigGlobal.PublicIngestBaseUrls, ",")),
		privateUploaders: createUploaders(strings.Split(ConfigGlobal.PrivateIngestBaseUrls, ",")),
	}

	if ConfigGlobal.EnableWebsockets {
		wsHub = newHub()
		go wsHub.run()
		go runHTTPServer()
	}
}

func createUploaders(targets []string) []uploader {
	var uploaders []uploader
	for _, target := range targets {
		if target == "" {
			continue
		}
		if len(target) < 4 {
			log.Infof("Got an ingest target that was less than 4 characters, not a valid ingest target: %v", target)
			continue
		}

		if target[0:8] == "http+pow" {
			uploaders = append(uploaders, newHTTPUploaderPow(target))
		} else if target[0:4] == "http" {
			uploaders = append(uploaders, newHTTPUploader(target))
		} else if target[0:4] == "nats" {
			uploaders = append(uploaders, newNATSUploader(target))
		} else {
			log.Infof("An invalid ingest target was specified: %v", target)
		}
	}

	return uploaders
}

func sendMsgToPublicUploaders(upload interface{}, topic string, state *albionState) {
	data, err := json.Marshal(upload)
	if err != nil {
		log.Errorf("Error while marshalling payload for %v: %v", err, topic)
		return
	}

	sendMsgToUploaders(data, topic, dis.publicUploaders)
	sendMsgToUploaders(data, topic, dis.privateUploaders)

	// If websockets are enabled, send the data there too
	if ConfigGlobal.EnableWebsockets {
		sendMsgToWebSockets(data, topic)
	}
}

func sendMsgToPrivateUploaders(upload lib.PersonalizedUpload, topic string, state *albionState) {
	if ConfigGlobal.DisableUpload {
		log.Info("Upload is disabled.")
		return
	}

	// TODO: Re-enable this when issue #14 is fixed
	// Will personalize with blanks for now in order to allow people to see the format
	// if state.CharacterName == "" || state.CharacterId == "" {
	// 	log.Error("The player name or id has not been set. Please restart the game and make sure the client is running.")
	// 	notification.Push("The player name or id has not been set. Please restart the game and make sure the client is running.")
	// 	return
	// }

	upload.Personalize(state.CharacterId, state.CharacterName)

	data, err := json.Marshal(upload)
	if err != nil {
		log.Errorf("Error while marshalling payload for %v: %v", err, topic)
		return
	}

	if len(dis.privateUploaders) > 0 {
		sendMsgToUploaders(data, topic, dis.privateUploaders)
	}

	// If websockets are enabled, send the data there too
	if ConfigGlobal.EnableWebsockets {
		sendMsgToWebSockets(data, topic)
	}
}

func sendMsgToUploaders(msg []byte, topic string, uploaders []uploader) {
	if ConfigGlobal.DisableUpload {
		log.Info("Upload is disabled.")
		return
	}

	for _, u := range uploaders {
		u.sendToIngest(msg, topic)
	}
}

func runHTTPServer() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(wsHub, w, r)
	})

	err := http.ListenAndServe(":8099", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sendMsgToWebSockets(msg []byte, topic string) {
	// TODO (gradius): send JSON data with topic string
	// TODO (gradius): this seems super hacky, and I'm sure there's a better way.
	var result string
	result = "{\"topic\": \"" + topic + "\", \"data\": " + string(msg) + "}"
	wsHub.broadcast <- []byte(result)
}
