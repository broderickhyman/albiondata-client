package client

import (
	"encoding/json"

	"github.com/regner/albionmarket-client/lib"
	"github.com/regner/albionmarket-client/log"
)

type operationAuctionGetRequestsResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

func (op operationAuctionGetRequestsResponse) Process(state *albionState, uploader iuploader) {
	log.Debug("Got response to AuctionGetOffers operation...")

	if state.LocationId == 0 {
		log.Warn("The players location has not yet been set. Please transition zones so the location can be identified.")
		return
	}

	orders := []*lib.MarketOrder{}

	for _, v := range op.MarketOrders {
		order := &lib.MarketOrder{}

		err := json.Unmarshal([]byte(v), order)
		if err != nil {
			log.Errorf("Problem converting market order to internal struct: %v", err)
		}

		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}

	log.Debugf("Sending %d market requests to ingest", len(orders))

	ingestRequest := lib.MarketUpload{
		Orders:     orders,
	}

	data, err := json.Marshal(ingestRequest)
	if err != nil {
		log.Errorf("Error while marshalling payload for market orders: %v", err)
		return
	}

	uploader.sendToIngest(data, "marketorders")
}
