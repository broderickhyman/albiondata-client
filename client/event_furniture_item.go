package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type eventFurnitureItem struct {
	IDInContainer int64  `mapstructure:"0"`
	ItemIndex     int64  `mapstructure:"1"`
	Quantity      int16  `mapstructure:"2"`
	Creator       string `mapstructure:"3"`
	Durability    int64  `mapstructure:"4"` // was multiplied by 10000
}

func (event eventFurnitureItem) Process(state *albionState) {
	log.Debug("Got equipment item event...")

	item := lib.ItemContainer{}
	item.AsFurnitureItem(event.ItemIndex, event.Quantity, event.Creator, event.Durability)

	state.ContainerItemsToSend[event.IDInContainer] = item
}
