package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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

// Prooves to the server that a pow was solved by submitting
// the pow and the solution as a POST request
func (u *httpUploaderPow) proovePow(pow Pow, solution string) {
	fullURL := u.baseURL + "/pow"

	resp, err := http.PostForm(fullURL, url.Values{
		"key": {pow.Key},
		"solution": {solution},
	})

	if err != nil {
		log.Errorf("Error while prooving pow: %v", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Errorf("HTTP Error while prooving pow. returned: %v", resp.StatusCode)
		return
	}
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
		hash := sha256.Sum256([]byte("aod^" + randhex + "^" + pow.Key))
		hexstring := fmt.Sprintf("%x", hash)
		bits := toBinaryBytes(hexstring)
		if strings.HasPrefix(bits, pow.Wanted) {
			log.Infof("SOLVED!")
			solution = randhex
			break
		}
	}
	return solution
}

func (u *httpUploaderPow) sendToIngest(body []byte, topic string) {
	client := &http.Client{Transport: u.transport}

	pow := Pow{}
	u.getPow(&pow)
	solution := solvePow(pow)
	u.proovePow(pow, solution)

	fullURL := u.baseURL + "/" + topic + "/" + pow.Key

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
