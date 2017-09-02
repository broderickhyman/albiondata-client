package client

import (
	"strconv"

	"github.com/regner/albiondata-client/log"
)

type operationJoinResponse struct {
	CharacterName      string `mapstructure:"2"`
	CharacterPartsJSON string `mapstructure:"6"`
	Location           string `mapstructure:"7"`
	Edition            string `mapstructure:"38"`
	GuildName          string `mapstructure:"47"`
}

func (op operationJoinResponse) Process(state *albionState) {
	log.Debugf("Got JoinResponse operation...")

	loc, err := strconv.Atoi(op.Location)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.Location)
		return
	}

	state.LocationId = loc
	log.Debugf("Updating player location to %v.", loc)
}
