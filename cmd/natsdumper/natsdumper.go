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
		fmt.Sprintf(
			"%s,%s,%s",
			lib.NatsMarketOrdersDeduped,
			lib.NatsGoldPricesDeduped,
			lib.NatsMapDataDeduped,
		),
		fmt.Sprintf(
			"NATS channels to connect to, comma saperated. Can be '%s', '%s', '%s', '%s'",
			lib.NatsMarketOrdersDeduped,
			lib.NatsGoldPricesDeduped,
			lib.NatsMarketOrdersIngest,
			lib.NatsGoldPricesIngest,
		),
	)
}

func subscribeMarketOrdersIngest(nc *nats.Conn) {
	fmt.Printf("mo i Subscribing %s\n", lib.NatsMarketOrdersIngest)
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
				fmt.Printf("mo i %s\n", string(jb))
			}
		}
	}
}

func subscribeMarketOrdersDeduped(nc *nats.Conn) {
	fmt.Printf("mo d Subscribing %s\n", lib.NatsMarketOrdersDeduped)
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
			fmt.Printf("mo d %s\n", string(msg.Data))
		}
	}
}

func subscribeGoldPricesIngest(nc *nats.Conn) {
	fmt.Printf("gp i Subscribing %s\n", lib.NatsGoldPricesIngest)
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
			fmt.Printf("gp i %s\n", string(msg.Data))
		}
	}
}

func subscribeGoldPricesDeduped(nc *nats.Conn) {
	fmt.Printf("gp d Subscribing %s\n", lib.NatsGoldPricesDeduped)
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
			fmt.Printf("gp d %s\n", string(msg.Data))
		}
	}
}

func subscribeMapDataIngest(nc *nats.Conn) {
	fmt.Printf("md i Subscribing %s\n", lib.NatsMapDataIngest)
	mapCh := make(chan *nats.Msg, 64)
	mapSub, err := nc.ChanSubscribe(lib.NatsMapDataIngest, mapCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer mapSub.Unsubscribe()

	for {
		select {
		case msg := <-mapCh:
			fmt.Printf("md i %s\n", string(msg.Data))
		}
	}
}

func subscribeMapDataDeduped(nc *nats.Conn) {
	fmt.Printf("md d Subscribing %s\n", lib.NatsMapDataDeduped)
	mapCh := make(chan *nats.Msg, 64)
	mapSub, err := nc.ChanSubscribe(lib.NatsMapDataDeduped, mapCh)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer mapSub.Unsubscribe()

	for {
		select {
		case msg := <-mapCh:
			fmt.Printf("md d %s\n", string(msg.Data))
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
			case lib.NatsMapDataIngest:
				go subscribeMapDataIngest(nc)
			case lib.NatsMapDataDeduped:
				go subscribeMapDataDeduped(nc)
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
	case lib.NatsMapDataIngest:
		subscribeMapDataIngest(nc)
	case lib.NatsMapDataDeduped:
		subscribeMapDataDeduped(nc)
	}
}
