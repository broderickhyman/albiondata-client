package client

import (
	"encoding/json"

	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
	"strconv"
)

type operationGetClusterMapInfo struct {
}

func (op operationGetClusterMapInfo) Process(state *albionState, uploader iuploader) {
	log.Debug("Got GetClusterMapInfo operation...")
}

type operationGetClusterMapInfoResponse struct {
	ZoneID          string   `mapstructure:"0"`
	BuildingType    []int    `mapstructure:"5"`
	AvailableFood   []int    `mapstructure:"10"`
	Reward          []int    `mapstructure:"12"`
	AvailableSilver []int    `mapstructure:"13"`
	Owners          []string `mapstructure:"14"`
	Buildable       []bool   `mapstructure:"19"`
	IsForSale       []bool   `mapstructure:"27"`
	BuyPrice        []int    `mapstructure:"28"`
}

func (op operationGetClusterMapInfoResponse) Process(state *albionState, uploader iuploader) {
	log.Debug("Got response to GetClusterMapInfo operation...")

	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.ZoneID)
		return
	}

	data, err := json.Marshal(lib.MapDataUpload{
		ZoneID:          zoneInt,
		BuildingType:    op.BuildingType,
		AvailableFood:   op.AvailableFood,
		Reward:          op.Reward,
		AvailableSilver: op.AvailableSilver,
		Owners:          op.Owners,
		Buildable:       op.Buildable,
		IsForSale:       op.IsForSale,
		BuyPrice:        op.BuyPrice,
	})

	if err != nil {
		log.Errorf("Error while marshalling payload for market data: %v", err)
		return
	}

	log.Info("Sending market data to ingest")
	uploader.sendToIngest(data, lib.NatsMapDataIngest)
}
