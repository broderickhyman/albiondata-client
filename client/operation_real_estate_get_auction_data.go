package client

import (
	"github.com/alexscott/albiondata-client/log"
)

type operationRealEstateGetAuctionData struct {
	PlotID int `mapstructure:"0"`
}

func (op operationRealEstateGetAuctionData) Process(state *albionState) {
	log.Debug("Got RealEstateGetAuctionData operation...")
}

type operationRealEstateGetAuctionDataResponse struct {
	Unknown           int    `mapstructure:"0"`
	HighestBidderName string `mapstructure:"1"`
	CurrentWinningBid int    `mapstructure:"2"`
	AuctionStartTime  int    `mapstructure:"3"`
	AuctionEndTime    int    `mapstructure:"4"`
}

func (op operationRealEstateGetAuctionDataResponse) Process(state *albionState) {
	log.Debug("Got response to RealEstateGetAuctionData operation...")
}
