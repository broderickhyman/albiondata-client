package operations

import (
	"log"

	"github.com/regner/albionmarket-client/client/albionstate"
)

type GoldMarketGetAverageInfo struct {
}

func (op GoldMarketGetAverageInfo) Process(state *albionstate.AlbionState) {
	log.Print("Got GoldMarketGetAverageInfo operation...")
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
	log.Print("Got response to GoldMarketGetAverageInfo operation...")
}
