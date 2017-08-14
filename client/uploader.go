package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/regner/albionmarket-client/log"
)

type uploader struct {
	transport *http.Transport
}

func newUploader() *uploader {
	return &uploader{
		transport: &http.Transport{},
	}
}

func (u *uploader) sendToIngest(body []byte, url string) {
	if ConfigGlobal.DisableUpload {
		return
	}

	client := &http.Client{Transport: u.transport}

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
		log.Errorf("Got bad response code: %v", resp.StatusCode)
		return
	}

	// See: https://stackoverflow.com/questions/17948827/reusing-http-connections-in-golang
	io.Copy(ioutil.Discard, resp.Body)

	log.Infof("Successfully sent ingest request to %v", ConfigGlobal.IngestBaseUrl)

	defer resp.Body.Close()
}
