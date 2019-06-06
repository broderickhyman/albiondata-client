package client

import (
	"albiondata-client/lib"
	"albiondata-client/log"
	"strconv"
)

type operationJoinResponse struct {
	CharacterID        lib.CharacterID `mapstructure:"1"`
	GuildID            lib.CharacterID `mapstructure:"45"`
	CharacterName      string          `mapstructure:"2"`
	CharacterPartsJSON string          `mapstructure:"6"`
	Location           string          `mapstructure:"7"`
	Edition            string          `mapstructure:"38"`
	GuildName          string          `mapstructure:"47"`
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

	state.CharacterId = op.CharacterID
	b64, _ := op.CharacterID.Base64()
	log.Debugf("Updating player ID to %v (%v).", op.CharacterID, b64)

	state.CharacterName = op.CharacterName
	log.Debugf("Updating player to %v.", op.CharacterName)
}
