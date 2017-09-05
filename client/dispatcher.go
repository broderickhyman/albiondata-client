package client

import (
	"net/http"

	"strings"

	"github.com/regner/albiondata-client/log"
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

	wsHub = newHub()
	go wsHub.run()
	go runHTTPServer()
}

func createUploaders(targets []string) []uploader {
	var uploaders []uploader

	for _, target := range targets {
		if len(target) < 4 {
			log.Infof("Got an ingest target that was less than 4 characters, not a valid ingest target: %v", target)
			continue
		}

		if target[0:4] == "http" {
			uploaders = append(uploaders, newHTTPUploader(target))
		} else if target[0:4] == "nats" {
			uploaders = append(uploaders, newNATSUploader(target))
		} else {
			log.Infof("An invalid ingest target was specified: %v", target)
		}
	}

	return uploaders
}

func sendMsgToPublicUploaders(msg []byte, topic string) {
	sendMsgToUploaders(msg, topic, dis.publicUploaders)
	sendMsgToUploaders(msg, topic, dis.privateUploaders)
	sendMsgToWebSockets(msg, topic)
}

func sendMsgToPrivateUploaders(msg []byte, topic string) {
	sendMsgToUploaders(msg, topic, dis.privateUploaders)
	sendMsgToWebSockets(msg, topic)
}

func sendMsgToUploaders(msg []byte, topic string, uploaders []uploader) {
	if ConfigGlobal.DisableUpload {
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
	var test string
	test = "{\"topic\": \"" + topic + "\", \"data\": " + string(msg) + "}"
	wsHub.broadcast <- []byte(test)
}
