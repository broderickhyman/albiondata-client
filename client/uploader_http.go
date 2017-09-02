package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/regner/albiondata-client/log"
)

type httpUploader struct {
	isPrivate bool
	baseURL   string
	transport *http.Transport
}

func newHTTPUploader(baseURL string, isPrivate bool) iuploader {
	return &httpUploader{
		isPrivate: isPrivate,
		baseURL:   baseURL,
		transport: &http.Transport{},
	}
}

func (u *httpUploader) private() bool {
	return u.isPrivate
}

func (u *httpUploader) sendToPrivateIngest(body []byte, queue string) {
	if u.private() {
		u.sendToIngest(body, queue, "PRIVATE")
	}
}

func (u *httpUploader) sendToPublicIngest(body []byte, queue string) {
	if u.private() {
		u.sendToIngest(body, queue, "PRIVATE")
	} else {
		u.sendToIngest(body, queue, "PUBLIC")
	}
}

func (u *httpUploader) sendToIngest(body []byte, queue string, privOrPublic string) {
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

	log.Debugf("Successfully sent %s ingest request to %v", privOrPublic, u.baseURL)

	defer resp.Body.Close()
}
