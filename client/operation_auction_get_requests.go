package client

import (
	"encoding/json"

	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
	"github.com/regner/albiondata-client/notification"
)

type operationAuctionGetRequestsResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

func (op operationAuctionGetRequestsResponse) Process(state *albionState) {
	log.Debug("Got response to AuctionGetOffers operation...")

	if state.LocationId == -1 {
		log.Warn("The players location has not yet been set. Please transition zones so the location can be identified.")
		notification.Push("The players location has not yet been set. Please transition zones so the location can be identified.")

		return
	}

	orders := []*lib.MarketOrder{}

	for _, v := range op.MarketOrders {
		order := &lib.MarketOrder{}

		err := json.Unmarshal([]byte(v), order)
		if err != nil {
			log.Errorf("Problem converting market order to internal struct: %v", err)
		}

		order.LocationID = state.LocationId
		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}

	upload := lib.MarketUpload{
		Orders: orders,
	}

	log.Infof("Sending %d market requests to ingest", len(orders))
	sendMsgToPublicUploaders(upload, lib.NatsMarketOrdersIngest, state)
}
