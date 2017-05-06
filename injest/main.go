package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"github.com/regner/amdr/shared"
	"os"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	r := gin.Default()

	r.POST("/api/v1/injest/", func(c *gin.Context) {
		var incomingRequest shared.InjestPostRequest
		c.BindJSON(&incomingRequest)

		var marketItems []shared.MarketItem
		for _, v := range incomingRequest.MarketItems {
			var item shared.MarketItem
			json.Unmarshal([]byte(v), &item)

			marketItems = append(marketItems, item)
		}

		natsMsg, _ := json.Marshal(marketItems)

		nc.Publish("amdr-injest", []byte(natsMsg))
	})

	r.Run(":8080")
}
