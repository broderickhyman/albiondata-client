package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	nats "github.com/nats-io/go-nats"
	"github.com/regner/albiondata-client/lib"
)

var (
	natsURL      string
	natsChannels string
)

func init() {
	flag.StringVar(
		&natsURL,
		"i",
		"nats://public:notsecure@ingest.albion-data.com:4222",
		"NATS URL to connect to.",
	)

	flag.StringVar(
		&natsChannels,
		"c",
		fmt.Sprintf("%s,%s", lib.NatsMarketOrdersDeduped, lib.NatsGoldPricesDeduped),
		fmt.Sprintf("NATS channels to connect to, comma saperated. Can be '%s', '%s', '%s', '%s'",
			lib.NatsMarketOrdersDeduped, lib.NatsGoldPricesDeduped, lib.NatsMarketOrdersIngest, lib.NatsGoldPricesIngest,
		),
	)
}

func subscribeMarketOrdersIngest(nc *nats.Conn) {
	fmt.Printf("mi Subscribing %s\n", lib.NatsMarketOrdersIngest)
	marketCh := make(chan *nats.Msg, 64)
	marketSub, err := nc.ChanSubscribe(lib.NatsMarketOrdersIngest, marketCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer marketSub.Unsubscribe()

	for {
		select {
		case msg := <-marketCh:
			morders := &lib.MarketUpload{}
			if err := json.Unmarshal(msg.Data, morders); err != nil {
				fmt.Printf("%v\n", err)
			}

			for _, order := range morders.Orders {
				jb, _ := json.Marshal(order)
				fmt.Printf("mi %s\n", string(jb))
			}
		}
	}
}

func subscribeMarketOrdersDeduped(nc *nats.Conn) {
	fmt.Printf("md Subscribing %s\n", lib.NatsMarketOrdersDeduped)
	marketCh := make(chan *nats.Msg, 64)
	marketSub, err := nc.ChanSubscribe(lib.NatsMarketOrdersDeduped, marketCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer marketSub.Unsubscribe()

	for {
		select {
		case msg := <-marketCh:
			fmt.Printf("md %s\n", string(msg.Data))
		}
	}
}

func subscribeGoldPricesIngest(nc *nats.Conn) {
	fmt.Printf("gi Subscribing %s\n", lib.NatsGoldPricesIngest)
	goldCh := make(chan *nats.Msg, 64)
	goldSub, err := nc.ChanSubscribe(lib.NatsGoldPricesIngest, goldCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer goldSub.Unsubscribe()

	for {
		select {
		case msg := <-goldCh:
			fmt.Printf("gi %s\n", string(msg.Data))
		}
	}
}

func subscribeGoldPricesDeduped(nc *nats.Conn) {
	fmt.Printf("gd Subscribing %s\n", lib.NatsGoldPricesDeduped)
	goldCh := make(chan *nats.Msg, 64)
	goldSub, err := nc.ChanSubscribe(lib.NatsGoldPricesDeduped, goldCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer goldSub.Unsubscribe()

	for {
		select {
		case msg := <-goldCh:
			fmt.Printf("gd %s\n", string(msg.Data))
		}
	}
}

func main() {
	flag.Parse()

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	chans := strings.Split(natsChannels, ",")

	if len(chans) > 1 {
		goChans := 0

		for _, goChan := range chans {
			switch goChan {
			case lib.NatsMarketOrdersIngest:
				go subscribeMarketOrdersIngest(nc)
			case lib.NatsMarketOrdersDeduped:
				go subscribeMarketOrdersDeduped(nc)
			case lib.NatsGoldPricesIngest:
				go subscribeGoldPricesIngest(nc)
			case lib.NatsGoldPricesDeduped:
				go subscribeGoldPricesDeduped(nc)
			}

			goChans = goChans + 1
			if goChans > len(chans)-2 {
				break
			}
		}
	}

	switch chans[len(chans)-1] {
	case lib.NatsMarketOrdersIngest:
		subscribeMarketOrdersIngest(nc)
	case lib.NatsMarketOrdersDeduped:
		subscribeMarketOrdersDeduped(nc)
	case lib.NatsGoldPricesIngest:
		subscribeGoldPricesIngest(nc)
	case lib.NatsGoldPricesDeduped:
		subscribeGoldPricesDeduped(nc)
	}
}
