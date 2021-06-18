package client

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/broderickhyman/albiondata-client/log"
)

type httpUploaderPow struct {
	baseURL   string
	transport *http.Transport
}

type Pow struct {
	Key    string `json:"key"`
	Wanted string `json:"wanted"`
}

// newHTTPUploaderPow creates a new HTTP uploader
func newHTTPUploaderPow(url string) uploader {
	return &httpUploaderPow{
		baseURL:   strings.Replace(url, "http+pow", "http", -1),
		transport: &http.Transport{},
	}
}

func (u *httpUploaderPow) getPow(target interface{}) {
	log.Debugf("GETTING POW")
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

// Prooves to the server that a pow was solved by submitting
// the pow's key, the solution and a nats msg as a POST request
// the topic becomes part of the URL
func (u *httpUploaderPow) uploadWithPow(pow Pow, solution string, natsmsg []byte, topic string) {

	fullURL := u.baseURL + "/pow/" + topic

	resp, err := http.PostForm(fullURL, url.Values{
		"key":      {pow.Key},
		"solution": {solution},
		"natsmsg":  {string(natsmsg)},
	})

	if err != nil {
		log.Errorf("Error while prooving pow: %v", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Errorf("HTTP Error while prooving pow. returned: %v", resp.StatusCode)
		return
	}

	log.Infof("Successfully sent ingest request to %v", u.baseURL)
}

// Generates a random hex string e.g.: faa2743d9181dca5
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Converts a string to bits e.g.: 0110011...
func toBinaryBytes(s string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(&buffer, "%08b", s[i])
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}

// Solves a pow looping through possible solutions
// until a correct one is found
// returns the solution
func solvePow(pow Pow) string {
	solution := ""
	for {
		randhex, _ := randomHex(8)
		if strings.HasPrefix(toBinaryBytes(fmt.Sprintf("%x", sha256.Sum256([]byte("aod^"+randhex+"^"+pow.Key)))), pow.Wanted) {
			log.Debugf("SOLVED!")
			solution = randhex
			break
		}
	}
	return solution
}

func (u *httpUploaderPow) sendToIngest(body []byte, topic string) {
	pow := Pow{}
	u.getPow(&pow)
	solution := solvePow(pow)
	u.uploadWithPow(pow, solution, body, topic)
}
