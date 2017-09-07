package client

import (
	"strconv"

	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

type operationGetClusterMapInfo struct {
}

func (op operationGetClusterMapInfo) Process(state *albionState) {
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

func (op operationGetClusterMapInfoResponse) Process(state *albionState) {
	log.Debug("Got response to GetClusterMapInfo operation...")

	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.ZoneID)
		return
	}

	upload := lib.MapDataUpload{
		ZoneID:          zoneInt,
		BuildingType:    op.BuildingType,
		AvailableFood:   op.AvailableFood,
		Reward:          op.Reward,
		AvailableSilver: op.AvailableSilver,
		Owners:          op.Owners,
		Buildable:       op.Buildable,
		IsForSale:       op.IsForSale,
		BuyPrice:        op.BuyPrice,
	}

	log.Info("Sending map data to ingest")
	sendMsgToPublicUploaders(upload, lib.NatsMapDataIngest, state)
}
