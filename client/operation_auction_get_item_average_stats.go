package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type operationAuctionGetItemAverageStats struct {
	ItemID      uint32        `mapstructure:"1"`
	Quality     uint8         `mapstructure:"2"`
	Timescale   lib.Timescale `mapstructure:"3"`
	Enchantment uint32        `mapstructure:"4"`
	MessageID   uint64        `mapstructure:"255"`
}

func (op operationAuctionGetItemAverageStats) Process(state *albionState) {
	var index = op.MessageID % CACHE_SIZE
	mhInfo := marketHistoryInfo{
		albionId:  op.ItemID,
		timescale: op.Timescale,
		quality:   op.Quality,
	}
	state.marketHistoryIDLookup[index] = mhInfo
	log.Debugf("Market History - Caching %d at %d.", mhInfo.albionId, index)
}

type operationAuctionGetItemAverageStatsResponse struct {
	ItemAmounts   []uint64 `mapstructure:"0"`
	SilverAmounts []uint64 `mapstructure:"1"`
	Timestamps    []uint64 `mapstructure:"2"`
	MessageID     int      `mapstructure:"255"`
}

func (op operationAuctionGetItemAverageStatsResponse) Process(state *albionState) {
	var index = op.MessageID % CACHE_SIZE
	var mhInfo = state.marketHistoryIDLookup[index]
	log.Debugf("Market History - Loaded itemID %d from cache at index %d", mhInfo.albionId, index)
	log.Debug("Got response to GetItemAverageStats operation for the itemID[", mhInfo.albionId, "] of quality: ", mhInfo.quality, " and on the timescale: ", mhInfo.timescale)

	if !state.IsValidLocation() {
		return
	}

	var histories []*lib.MarketHistory

	// TODO can we make this safer? Right now we just assume all the arrays are the same length as the number of item amounts
	for i := range op.ItemAmounts {

		history := &lib.MarketHistory{}
		history.ItemAmount = op.ItemAmounts[i]
		history.SilverAmount = op.SilverAmounts[i]
		history.Timestamp = op.Timestamps[i]
		histories = append(histories, history)
	}

	if len(histories) < 1 {
		log.Info("Auction Stats Response - no history\n\n")
		return
	}

	upload := lib.MarketHistoriesUpload{
		AlbionID:     mhInfo.albionId,
		LocationID:   state.LocationId,
		QualityLevel: mhInfo.quality,
		Timescale:    mhInfo.timescale,
		Histories:    histories,
	}

	log.Infof("Sending %d item average stats to ingest for albionID %d", len(histories), mhInfo.albionId)
	sendMsgToPublicUploaders(upload, lib.NatsMarketHistoriesIngest, state)
}
