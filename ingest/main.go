package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"github.com/regner/albion-market-data-relay/shared"
	"os"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	r := gin.Default()

	r.POST("/api/v1/ingest/", func(c *gin.Context) {
		var incomingRequest shared.InjestPostRequest
		c.BindJSON(&incomingRequest)

		var marketItems []shared.MarketItem
		for _, v := range incomingRequest.MarketItems {
			var item shared.MarketItem
			json.Unmarshal([]byte(v), &item)

			marketItems = append(marketItems, item)
		}

		natsMsg, _ := json.Marshal(marketItems)

		nc.Publish("amdr-ingest", []byte(natsMsg))
	})

	r.Run(":8080")
}
