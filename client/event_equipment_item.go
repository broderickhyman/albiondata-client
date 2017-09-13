package client

import (
	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

type eventEquipmentItem struct {
	IDInContainer int64  `mapstructure:"0"`
	ItemIndex     int64  `mapstructure:"1"`
	Quantity      int16  `mapstructure:"2"`
	Creator       string `mapstructure:"3"`
	Quality       byte   `mapstructure:"4"`
	Durability    int64  `mapstructure:"5"` // was multiplied by 10000
}

func (event eventEquipmentItem) Process(state *albionState) {
	log.Debug("Got equipment item event...")

	item := lib.ItemContainer{}
	item.AsEquipmentItem(event.ItemIndex, event.Quantity, event.Creator, event.Quality, event.Durability)

	state.ContainerItemsToSend[event.IDInContainer] = item
}
