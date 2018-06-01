package client

import (
	"strconv"

	"github.com/broderickhyman/albiondata-client/log"
)

type operationGetGameServerByCluster struct {
	ZoneID string `mapstructure:"0"`
}

func (op operationGetGameServerByCluster) Process(state *albionState) {
	log.Debug("Got GetGameServerByCluster operation...")

	state.LocationString = op.ZoneID
	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.ZoneID)
		state.LocationId = -2 // hack
		return
	}

	log.Debugf("Updating player location to %v.", zoneInt)
	state.LocationId = zoneInt
}
