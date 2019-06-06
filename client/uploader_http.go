package client

import (
	"albiondata-client/log"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type httpUploader struct {
	baseURL   string
	transport *http.Transport
}

// newHTTPUploader creates a new HTTP uploader
func newHTTPUploader(url string) uploader {
	return &httpUploader{
		baseURL:   url,
		transport: &http.Transport{},
	}
}

func (u *httpUploader) sendToIngest(body []byte, topic string) {
	client := &http.Client{Transport: u.transport}

	fullURL := u.baseURL + "/" + topic

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

	log.Infof("Successfully sent ingest request to %v", u.baseURL)

	defer resp.Body.Close()
}
