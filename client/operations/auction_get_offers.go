package operations

import (
	"log"

	"github.com/regner/albionmarket-client/client/albionstate"
)

type AuctionGetOffers struct {
	Category         string   `mapstructure:"1"`
	SubCategory      string   `mapstructure:"2"`
	Quality          string   `mapstructure:"3"`
	Enchantment      uint32   `mapstructure:"4"`
	EnchantmentLevel string   `mapstructure:"8"`
	Tier             string   `mapstructure:"5"`
	ItemIds          []uint16 `mapstructure:"6"`
	MaxResults       uint32   `mapstructure:"9"`
	IsAscendingOrder bool     `mapstructure:"11"`
}

func (op AuctionGetOffers) Process(state *albionstate.AlbionState) {
	log.Print("Got AuctionGetOffers operation...")
}
