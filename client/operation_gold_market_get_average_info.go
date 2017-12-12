package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type operationGoldMarketGetAverageInfo struct {
}

func (op operationGoldMarketGetAverageInfo) Process(state *albionState) {
	log.Debug("Got GoldMarketGetAverageInfo operation...")
}

type operationGoldMarketGetAverageInfoResponse struct {
	GoldPrices []int   `mapstructure:"0"`
	TimeStamps []int64 `mapstructure:"1"`
}

func (op operationGoldMarketGetAverageInfoResponse) Process(state *albionState) {
	log.Debug("Got response to GoldMarketGetAverageInfo operation...")

	upload := lib.GoldPricesUpload{
		Prices:     op.GoldPrices,
		TimeStamps: op.TimeStamps,
	}

	log.Info("Sending gold prices to ingest")
	sendMsgToPublicUploaders(upload, lib.NatsGoldPricesIngest, state)
}
