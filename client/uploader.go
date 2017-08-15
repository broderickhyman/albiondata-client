package client

import (
	"bytes"
	"net/http"

	"github.com/regner/albionmarket-client/log"
)

func uploaderSendToIngest(body []byte, url string) {
	if ConfigGlobal.DisableUpload {
		return
	}

	client := &http.Client{}

	fullUrl := ConfigGlobal.IngestBaseUrl + url

	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Errorf("Error while create new request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error while sending ingest with data: %v", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Errorf("Got bad response code (%v) when uploading to: %v", resp.StatusCode, fullUrl)
		return
	}

	log.Infof("Successfully sent ingest request to %v", ConfigGlobal.IngestBaseUrl)

	defer resp.Body.Close()
}
