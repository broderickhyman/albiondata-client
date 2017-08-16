package client

import (
	"encoding/json"

	"github.com/regner/albionmarket-client/log"
)

type operationAuctionGetOffers struct {
	Category         string   `mapstructure:"1"`
	SubCategory      string   `mapstructure:"2"`
	Quality          string   `mapstructure:"3"`
	Enchantment      uint32   `mapstructure:"4"`
	EnchantmentLevel string   `mapstructure:"8"`
	ItemIds          []uint16 `mapstructure:"6"`
	MaxResults       uint32   `mapstructure:"9"`
	IsAscendingOrder bool     `mapstructure:"11"`
}

func (op operationAuctionGetOffers) Process(state *albionState, uploader iuploader) {
	log.Debug("Got AuctionGetOffers operation...")
}

type operationAuctionGetOffersResponse struct {
	MarketOrders []string `mapstructure:"0"`
}

// MarketOrder contains an order (offer or request)
type MarketOrder struct {
	ID               int    `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
	LocationID       int    `json:"LocationId"`
	QualityLevel     int    `json:"QualityLevel"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	Price            int    `json:"UnitPriceSilver"`
	Amount           int    `json:"Amount"`
	AuctionType      string `json:"AuctionType"`
	Expires          string `json:"Expires"`
}

// MarketUpload contains a list of orders and the location where the orders are from
type MarketUpload struct {
	Orders     []*MarketOrder
	LocationID int
}

func (op operationAuctionGetOffersResponse) Process(state *albionState, uploader iuploader) {
	log.Debug("Got response to AuctionGetOffers operation...")

	if state.LocationId == 0 {
		log.Error("The players location has not yet been set. Please transition zones so the location can be identified.")
		return
	}

	orders := []*MarketOrder{}

	for _, v := range op.MarketOrders {
		order := &MarketOrder{}

		err := json.Unmarshal([]byte(v), order)
		if err != nil {
			log.Errorf("Problem converting market order to internal struct: %v", err)
		}

		orders = append(orders, order)
	}

	if len(orders) < 1 {
		return
	}

	log.Debugf("Sending %d market offers to ingest", len(orders))

	ingestRequest := MarketUpload{
		Orders:     orders,
		LocationID: state.LocationId,
	}

	data, err := json.Marshal(ingestRequest)
	if err != nil {
		log.Errorf("Error while marshalling payload for market orders: %v", err)
		return
	}

	uploader.sendToIngest(data, "marketorders")
}
