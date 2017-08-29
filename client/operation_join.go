package client

import (
	"github.com/regner/albiondata-client/log"
)

type operationJoinResponse struct {
	CharacterName      string `mapstructure:"2"`
	CharacterPartsJSON string `mapstructure:"6"`
	Location           string `mapstructure:"7"`
	Edition            string `mapstructure:"38"`
	GuildName          string `mapstructure:"47"`
}

func (op operationJoinResponse) Process(state *albionState, uploader iuploader) {
	log.Debugf("Got JoinResponse operation...")
}
