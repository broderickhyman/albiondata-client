package operations

import (
	"log"
	"strconv"
	"strings"

	"github.com/regner/albionmarket-client/client/albionstate"
)

//10
type GetGameServerByCluster struct {
	ZoneID               string `mapstructure:"0"`
	NrZoneChangesSession int32  `mapstructure:"255"`
}

func (op GetGameServerByCluster) Process(state *albionstate.AlbionState) {
	// By having a "." it means this zone is an instance, such as an island.
	if strings.Contains(op.ZoneID, ".") {
		return
	}

	zoneInt, err := strconv.Atoi(op.ZoneID)
	if err != nil {
		log.Printf("Unable to convert zoneID to int. This is bad. ZoneID: %v", op.ZoneID)
	}

	log.Printf("Updating player location to %v.", zoneInt)
	state.LocationId = zoneInt
}
