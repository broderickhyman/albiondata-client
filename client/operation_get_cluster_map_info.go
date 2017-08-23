package client

import (
	"github.com/regner/albiondata-client/log"
)

type operationGetClusterMapInfo struct {
}

func (op operationGetClusterMapInfo) Process(state *albionState, uploader iuploader) {
	log.Debug("Got GetClusterMapInfo operation...")
}

type operationGetClusterMapInfoResponse struct {
	ZoneID          int      `mapstructure:"0"`
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
}
