package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"

	"github.com/broderickhyman/albiondata-client/log"
)

type httpUploaderPow struct {
	baseURL   string
	transport *http.Transport
}

type Pow struct {
	Key  string  `json:"key"`
	Wanted  string  `json:"wanted"`
}

// newHTTPUploaderPow creates a new HTTP uploader
func newHTTPUploaderPow(url string) uploader {
	return &httpUploaderPow{
		baseURL: strings.Replace(url, "http+pow", "http", -1),
		transport: &http.Transport{},
	}
}

func (u *httpUploaderPow) getPow(target interface{}) {
	log.Infof("GETTING POW")
	fullURL := u.baseURL + "/pow"

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Errorf("Error in Pow Get request: %v", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Errorf("Got bad response code: %v", resp.StatusCode)
		return
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Errorf("Error in parsing Pow Get request: %v", err)
		return
	}
}

func (u *httpUploaderPow) sendToIngest(body []byte, topic string) {
	client := &http.Client{Transport: u.transport}

	fullURL := u.baseURL + "/" + topic

	pow := Pow{}
	u.getPow(&pow)
	log.Infof("POW: %v, %v", pow.Key, pow.Wanted,)

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
