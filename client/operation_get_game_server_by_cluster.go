package client

import (
	"strconv"
	"strings"

	"github.com/broderickhyman/albiondata-client/log"
)

type operationGetGameServerByCluster struct {
	ZoneID string `mapstructure:"0"`
}

func (op operationGetGameServerByCluster) Process(state *albionState) {
	log.Debug("Got GetGameServerByCluster operation...")

	state.LocationString = op.ZoneID
	// TODO: Fix hack for second caerleon marketplace
	// Most likely will need to only use strings for player location in client
	zoneInt, err := strconv.Atoi(strings.ReplaceAll(op.ZoneID, "-Auction2", ""))
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.ZoneID)
		state.LocationId = -2 // hack
		return
	}

	log.Infof("Updating player location to %v.", zoneInt)
	state.LocationId = zoneInt
}
