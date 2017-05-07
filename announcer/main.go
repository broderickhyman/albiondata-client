package main

import (
	"os"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/go-nats"
	"github.com/gorilla/websocket"
	"github.com/regner/albion-market-data-relay/shared"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func main() {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync(shared.NatsTopic)

	r := gin.Default()

	r.GET("/api/v1/announce/", func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Failed to set websocket upgrade: %+v", err)
			return
		}

		for {
			m, _ := sub.NextMsg(1 * time.Second)
			fmt.Println(m)
			conn.WriteMessage(websocket.TextMessage, []byte("test"))
		}
	})

	r.Run(":8081")
}