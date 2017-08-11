package operations

import (
	"encoding/json"
	"log"

	"github.com/regner/albionmarket-client/client/albionstate"
	"github.com/regner/albionmarket-client/client/config"
	"github.com/regner/albionmarket-client/client/uploader"
)

type AuctionGetRequestsResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

func (op AuctionGetRequestsResponse) Process(state *albionstate.AlbionState) {
	log.Print("Got response to AuctionGetOffers operation...")

	if state.LocationId == 0 {
		log.Printf("The players location has not yet been set. Pleas transition zones so the location can be identified.")
		return
	}

	orders := []*marketOrder{}

	for _, v := range op.MarketOrders {
		order := &marketOrder{}

		err := json.Unmarshal([]byte(v), order)
		if err != nil {
			log.Printf("Problem converting market order to internal struct: %v", err)
		}

		orders = append(orders, order)
	}

	if len(orders) > 0 {
		ingestRequest := marketUpload{
			Orders:     orders,
			LocationID: state.LocationId,
		}

		data, err := json.Marshal(ingestRequest)
		if err != nil {
			log.Printf("Error while marshalling payload for market orders: %v", err)
			return
		}

		if !config.GlobalConfiguration.DisableUpload {
			uploader.SendToIngest([]byte(string(data)), config.GlobalConfiguration.IngestUrl)
		}
	}
}
