package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type eventJournalItem struct {
	IDInContainer int64  `mapstructure:"0"`
	ItemIndex     int64  `mapstructure:"1"`
	Quantity      int16  `mapstructure:"2"`
	Creator       string `mapstructure:"3"`
	FameStored    int64  `mapstructure:"6"` // was multiplied by 10000
}

func (event eventJournalItem) Process(state *albionState) {
	log.Debug("Got journal item event...")

	item := lib.ItemContainer{}
	item.AsJournalItem(event.ItemIndex, event.Quantity, event.Creator, event.FameStored)

	state.ContainerItemsToSend[event.IDInContainer] = item
}
