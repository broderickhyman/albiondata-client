package client

import (
	"sort"

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
	var index = op.MessageID % CacheSize
	mhInfo := marketHistoryInfo{
		albionId:  op.ItemID,
		timescale: op.Timescale,
		quality:   op.Quality,
	}
	state.marketHistoryIDLookup[index] = mhInfo
	log.Debugf("Market History - Caching %d at %d.", mhInfo.albionId, index)
}

type operationAuctionGetItemAverageStatsResponse struct {
	ItemAmounts   []int64  `mapstructure:"0"`
	SilverAmounts []uint64 `mapstructure:"1"`
	Timestamps    []uint64 `mapstructure:"2"`
	MessageID     int      `mapstructure:"255"`
}

func (op operationAuctionGetItemAverageStatsResponse) Process(state *albionState) {
	var index = op.MessageID % CacheSize
	var mhInfo = state.marketHistoryIDLookup[index]
	log.Debugf("Market History - Loaded itemID %d from cache at index %d", mhInfo.albionId, index)
	log.Debug("Got response to GetItemAverageStats operation for the itemID[", mhInfo.albionId, "] of quality: ", mhInfo.quality, " and on the timescale: ", mhInfo.timescale)

	if !state.IsValidLocation() {
		return
	}

	var histories []*lib.MarketHistory

	// TODO can we make this safer? Right now we just assume all the arrays are the same length as the number of item amounts
	for i := range op.ItemAmounts {
		// sometimes opAuctionGetItemAverageStats receives negative item amounts
		if op.ItemAmounts[i] < 0 {
			if op.ItemAmounts[i] < -124 {
				// still don't know what to do with these
				log.Debugf("Market History - Ignoring negative item amount %d for %d silver on %d", op.ItemAmounts[i], op.SilverAmounts[i], op.Timestamps[i])
				continue
			}
			// however these can be interpreted by adding them to 256
			// TODO: make more sense of this, (perhaps there is a better way)
			log.Debugf("Market History - Interpreting negative item amount %d as %d for %d silver on %d", op.ItemAmounts[i], 256+op.ItemAmounts[i], op.SilverAmounts[i], op.Timestamps[i])
			op.ItemAmounts[i] = 256 + op.ItemAmounts[i]
		}
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

	// Sort history by descending time so the newest is always first in the list
	sort.SliceStable(histories, func(i, j int) bool {
		return histories[i].Timestamp > histories[j].Timestamp
	})

	upload := lib.MarketHistoriesUpload{
		AlbionId:     mhInfo.albionId,
		LocationId:   state.LocationId,
		QualityLevel: mhInfo.quality,
		Timescale:    mhInfo.timescale,
		Histories:    histories,
	}

	log.Infof("Sending %d item average stats to ingest for albionID %d", len(histories), mhInfo.albionId)
	sendMsgToPublicUploaders(upload, lib.NatsMarketHistoriesIngest, state)
}
