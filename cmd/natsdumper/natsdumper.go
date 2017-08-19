package main

import (
	"encoding/json"
	"flag"
	"fmt"

	nats "github.com/nats-io/go-nats"
	"github.com/regner/albiondata-client/lib"
)

var natsURL string

func init() {
	flag.StringVar(
		&natsURL,
		"i",
		"nats://public:notsecure@ingest.albion-data.com:4222",
		"NATS URL to connect to.",
	)
}

func dumpMarketOrders(m *nats.Msg) {
	morders := &lib.MarketUpload{}
	if err := json.Unmarshal(m.Data, morders); err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, order := range morders.Orders {
		jb, _ := json.Marshal(order)
		fmt.Printf("%d %s\n", order.LocationID, string(jb))
	}
}

func main() {
	flag.Parse()

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	marketCh := make(chan *nats.Msg, 64)
	marketSub, err := nc.ChanSubscribe(lib.NatsGoldPricesDeduped, marketCh)
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
