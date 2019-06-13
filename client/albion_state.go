package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
)

type albionState struct {
	LocationId     int
	LocationString string
	CharacterId    lib.CharacterID
	CharacterName  string
}
