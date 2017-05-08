package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type InjestRequest struct {
	MarketItems []string
}

var injestUrl = "http://192.168.1.20:8080/api/v1/ingest/"

func SendMarketItems(marketItems []string) {
	client := &http.Client{}

	injestRequest := InjestRequest{
		MarketItems: marketItems,
	}

	data, err := json.Marshal(injestRequest)
	if err != nil {
		log.Printf("Error while marshalling payload: %v", err)
	}

	req, err := http.NewRequest("POST", injestUrl, bytes.NewBuffer([]byte(string(data))))
	if err != nil {
		log.Printf("Error while create new reqest: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error while sending market data: %v", err)
	}

	defer resp.Body.Close()
}
