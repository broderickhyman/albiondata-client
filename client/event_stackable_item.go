package client

import (
	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

type eventStackableItem struct {
	IDInContainer int64 `mapstructure:"0"`
	ItemIndex     int64 `mapstructure:"1"`
	Quantity      int16 `mapstructure:"2"`
}

func (event eventStackableItem) Process(state *albionState) {
	log.Debug("Got stackable item event...")

	item := lib.ItemContainer{}
	item.AsStackableItem(event.ItemIndex, event.Quantity)

	state.ContainerItemsToSend[event.IDInContainer] = item
}
