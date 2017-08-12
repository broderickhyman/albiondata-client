package uploader

import (
	"bytes"
	"net/http"

	"github.com/regner/albionmarket-client/client/config"
	"github.com/regner/albionmarket-client/log"
)

func SendToIngest(body []byte, url string) {
	if config.GlobalConfiguration.DisableUpload {
		return
	}

	client := &http.Client{}

	fullUrl := config.GlobalConfiguration.IngestBaseUrl + url

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
		log.Errorf("Got bad response code: %v", resp.StatusCode)
		return
	}

	log.Infof("Successfully sent ingest request to %v", config.GlobalConfiguration.IngestBaseUrl)

	defer resp.Body.Close()
}
