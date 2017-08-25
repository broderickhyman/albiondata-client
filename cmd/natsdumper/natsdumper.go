package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	nats "github.com/nats-io/go-nats"
	"github.com/regner/albiondata-client/lib"
)

var (
	natsURL         string
	natsChannels    string
	saveLocally     bool
	saveLocallyPath string
	csvPerMessage   bool
	timestamp       string = "_" + time.Now().Format("20060102150405")
)

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func writeDataToFile(msg []string, identification string, filename string) bool {
	fExists := fileExists(filename)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	if !fExists {
		var jsonKeys []string
		switch identification {
		case lib.NatsGoldPricesDeduped, lib.NatsGoldPricesIngest:
			jsonKeys = lib.GetGoldPricesUploadJsonKeys()
		case lib.NatsMarketOrdersDeduped, lib.NatsMarketOrdersIngest:
			jsonKeys = lib.GetMarketOrderJsonKeys()
		case lib.NatsMapDataDeduped, lib.NatsMapDataIngest:
			jsonKeys = lib.MapDataHeaders()
		}

		csvWriter.Write(jsonKeys)
	}

	csvWriter.Write(msg)
	csvWriter.Flush()

	return true
}

func saveToCSVFile(msg []string, identification string) {
	if csvPerMessage {
		timestamp = "_" + time.Now().Format("20060102150405")
	}
	var filename string = filepath.Join(saveLocallyPath, identification+timestamp+".csv")

	absFilename, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	os.MkdirAll(filepath.Dir(absFilename), os.ModePerm)

	if !writeDataToFile(msg, identification, absFilename) {
		fmt.Printf("Failed to write to file %v\n", absFilename)
		return
	}
}

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

	flag.StringVar(
		&saveLocallyPath,
		"s",
		"",
		"If specified all uploads will be saved locally to this folder.",
	)

	flag.BoolVar(
		&csvPerMessage,
		"nsf",
		false,
		"If specified a new csv file will be created for every incoming message.",
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
				if saveLocally {
					saveToCSVFile(order.StringArray(), lib.NatsMarketOrdersIngest)
				}
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
			if saveLocally {
				order := &lib.MarketOrder{}
				if err := json.Unmarshal(msg.Data, order); err != nil {
					fmt.Printf("%v\n", err)
				}
				saveToCSVFile(order.StringArray(), lib.NatsMarketOrdersDeduped)
			}
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
			if saveLocally {
				gp := &lib.GoldPricesUpload{}
				if err := json.Unmarshal(msg.Data, gp); err != nil {
					fmt.Printf("%v\n", err)
				}
				sas := gp.StringArrays()
				for _, sa := range sas {
					saveToCSVFile(sa, lib.NatsGoldPricesIngest)
				}
			}
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
			if saveLocally {
				gp := &lib.GoldPricesUpload{}
				if err := json.Unmarshal(msg.Data, gp); err != nil {
					fmt.Printf("%v\n", err)
				}
				sas := gp.StringArrays()
				for _, sa := range sas {
					saveToCSVFile(sa, lib.NatsGoldPricesIngest)
				}
			}
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
			if saveLocally {
				md := &lib.MapDataUpload{}
				if err := json.Unmarshal(msg.Data, md); err != nil {
					fmt.Printf("%v\n", err)
				}
				sas := md.StringArrays()
				for _, sa := range sas {
					saveToCSVFile(sa, lib.NatsMapDataIngest)
				}
			}
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
			if saveLocally {
				md := &lib.MapDataUpload{}
				if err := json.Unmarshal(msg.Data, md); err != nil {
					fmt.Printf("%v\n", err)
				}
				sas := md.StringArrays()
				for _, sa := range sas {
					saveToCSVFile(sa, lib.NatsMapDataIngest)
				}
			}
		}
	}
}

func main() {
	flag.Parse()

	if saveLocallyPath != "" {
		saveLocally = true
	}

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
