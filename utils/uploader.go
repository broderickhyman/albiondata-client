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

func SendMarketItems(marketItems []string, ingestUrl string) {
	client := &http.Client{}

	injestRequest := InjestRequest{
		MarketItems: marketItems,
	}

	data, err := json.Marshal(injestRequest)
	if err != nil {
		log.Printf("Error while marshalling payload: %v", err)
		return
	}

	req, err := http.NewRequest("POST", ingestUrl, bytes.NewBuffer([]byte(string(data))))
	if err != nil {
		log.Printf("Error while create new reqest: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error while sending market data: %v", err)
		return
	}

	defer resp.Body.Close()
}
