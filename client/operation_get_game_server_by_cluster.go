package client

import (
	"strconv"
	"strings"

	"github.com/regner/albionmarket-client/log"
)

type operationGetGameServerByCluster struct {
	ZoneID               string `mapstructure:"0"`
	NrZoneChangesSession int32  `mapstructure:"255"`
}

func (op operationGetGameServerByCluster) Process(state *albionState, uploader *uploader) {
	log.Debug("Got GetGameServerByCluster operation...")

	// By having a "." it means this zone is an instance, such as an island.
	if strings.Contains(op.ZoneID, ".") {
		return
	}

	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Errorf("Unable to convert zoneID to int. This is bad. ZoneID: %v", op.ZoneID)
	}

	log.Debugf("Updating player location to %v.", zoneInt)
	state.LocationId = zoneInt
}
