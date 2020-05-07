package client

import (
	"strconv"

	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
)

type operationJoinResponse struct {
	CharacterID   lib.CharacterID `mapstructure:"1"`
	CharacterName string          `mapstructure:"2"`
	Location      string          `mapstructure:"8"`
	GuildID       lib.CharacterID `mapstructure:"47"`
	GuildName     string          `mapstructure:"51"`
}

//CharacterPartsJSON string          `mapstructure:"6"`
//Edition            string          `mapstructure:"38"`

func (op operationJoinResponse) Process(state *albionState) {
	log.Debugf("Got JoinResponse operation...")

	loc, err := strconv.Atoi(op.Location)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.")
		state.LocationId = -2
	} else {
		state.LocationId = loc
	}
	log.Infof("Updating player location to %v.", op.Location)

	state.CharacterId = op.CharacterID
	log.Infof("Updating player ID to %v.", op.CharacterID)

	state.CharacterName = op.CharacterName
	log.Infof("Updating player to %v.", op.CharacterName)
}
