package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nats-io/go-nats"
	"github.com/regner/albion-market-data-relay/shared"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	registry := NewRegistry()
	go registry.run()

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	ec.Subscribe(shared.NatsTopic, func(marketUpdate string) {
		registry.broadcast <- []byte(marketUpdate)
	})

	r := gin.Default()

	r.GET("/api/v1/announce/", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := &Client{registry: registry, conn: conn, send: make(chan []byte, 256)}
		client.registry.register <- client

		go client.run()
	})

	r.Run(":8081")
}
