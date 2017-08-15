package main

import (
	"encoding/json"
	"flag"
	"fmt"

	nats "github.com/nats-io/go-nats"
	"github.com/regner/albionmarket-client/client"
)

var natsURL string

func init() {
	flag.StringVar(
		&natsURL,
		"i",
		"nats://localhost:2222",
		"NATS URL to subscribe to.",
	)
}

func dumpMarketOrders(m *nats.Msg) {
	morders := &client.MarketUpload{}
	if err := json.Unmarshal(m.Data, morders); err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, order := range morders.Orders {
		jb, _ := json.Marshal(order)
		fmt.Printf("%d %s\n", morders.LocationID, string(jb))
	}
}

func main() {
	flag.Parse()

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	marketCh := make(chan *nats.Msg, 64)
	marketSub, err := nc.ChanSubscribe("marketorders", marketCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer marketSub.Unsubscribe()

	for {
		select {
		case msg := <-marketCh:
			dumpMarketOrders(msg)
		}
	}
}
