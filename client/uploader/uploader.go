package uploader

import (
	"bytes"
	"github.com/regner/albionmarket-client/log"
	"net/http"

	"github.com/regner/albionmarket-client/client/config"
)

func SendToIngest(body []byte, url string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error while create new reqest: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error while sending ingest with data: %v", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Printf("Got bad response code: %v", resp.StatusCode)
		return
	}

	log.Printf("Sucesfully sent ingest request to %v", config.GlobalConfiguration.IngestUrl)

	defer resp.Body.Close()
}
