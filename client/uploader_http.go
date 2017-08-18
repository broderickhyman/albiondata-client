package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/regner/albiondata-client/log"
)

type httpUploader struct {
	baseURL   string
	transport *http.Transport
}

func newHTTPUploader(baseURL string) iuploader {
	return &httpUploader{
		baseURL:   baseURL,
		transport: &http.Transport{},
	}
}

func (u *httpUploader) sendToIngest(body []byte, queue string) {
	client := &http.Client{Transport: u.transport}

	fullURL := u.baseURL + "/" + queue

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer([]byte(body)))
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
