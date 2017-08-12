package operations

import (
	"encoding/json"

	"github.com/regner/albionmarket-client/client/albionstate"
	"github.com/regner/albionmarket-client/client/uploader"
	"github.com/regner/albionmarket-client/log"
)

type GoldMarketGetAverageInfo struct {
}

func (op GoldMarketGetAverageInfo) Process(state *albionstate.AlbionState) {
	log.Debug("Got GoldMarketGetAverageInfo operation...")
}

type GoldMarketGetAverageInfoResponse struct {
	GoldPrices []int `mapstructure:"0"`
	TimeStamps []int `mapstructure:"1"`
}

type goldInfoUpload struct {
	Prices     []int
	TimeStamps []int
}

func (op GoldMarketGetAverageInfoResponse) Process(state *albionstate.AlbionState) {
	log.Debug("Got response to GoldMarketGetAverageInfo operation...")

	data, err := json.Marshal(op)
	if err != nil {
		log.Errorf("Error while marshalling payload for gold prices: %v", err)
		return
	}

	uploader.SendToIngest([]byte(string(data)), "goldprices")
}
