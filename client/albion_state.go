package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
	"github.com/broderickhyman/albiondata-client/notification"
)

type albionState struct {
	LocationId     int
	LocationString string
	CharacterId    lib.CharacterID
	CharacterName  string
}

func (state albionState) IsValidLocation() bool {
	if state.LocationId < 0 {
		if state.LocationId == -1 {
			log.Error("The players location has not yet been set. Please transition zones so the location can be identified.")
			notification.Push("The players location has not yet been set. Please transition zones so the location can be identified.")
		} else {
			log.Error("The players location is not valid. Please transition zones so the location can be fixed.")
			notification.Push("The players location is not valid. Please transition zones so the location can be fixed.")
		}
		return false
	}
	return true
}
