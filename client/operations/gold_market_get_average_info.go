package operations

import (
	"github.com/regner/albionmarket-client/log"

	"github.com/regner/albionmarket-client/client/albionstate"
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
}
