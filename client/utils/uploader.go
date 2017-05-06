package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		fmt.Println("SOMETHING BAD HAPPENED!")
	}

	req, err := http.NewRequest("POST", injestUrl, bytes.NewBuffer([]byte(string(data))))
	if err != nil {
		fmt.Println("SOMETHING BAD HAPPENED!")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("SOMETHING BAD HAPPENED!")
	}

	defer resp.Body.Close()
}
