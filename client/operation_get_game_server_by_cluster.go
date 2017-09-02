package client

import (
	"strconv"

	"github.com/regner/albiondata-client/log"
)

type operationGetGameServerByCluster struct {
	ZoneID string `mapstructure:"0"`
}

func (op operationGetGameServerByCluster) Process(state *albionState) {
	log.Debug("Got GetGameServerByCluster operation...")

	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.ZoneID)
		return
	}

	log.Debugf("Updating player location to %v.", zoneInt)
	state.LocationId = zoneInt
}
