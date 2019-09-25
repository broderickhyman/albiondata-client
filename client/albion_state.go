package client

import (
	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
	"github.com/broderickhyman/albiondata-client/notification"
)

const CACHE_SIZE = 256

type marketHistoryInfo struct {
	albionId  uint32
	timescale lib.Timescale
	quality   uint8
}

type albionState struct {
	LocationId     int
	LocationString string
	CharacterId    lib.CharacterID
	CharacterName  string

	// A lot of information is sent out but not contained in the response when requesting marketHistory (e.g. ID)
	// This information is stored in marketHistoryInfo
	// This array acts as a type of cache for that info
	// The index is the message number (param255) % CACHE_SIZE
	marketHistoryIDLookup [CACHE_SIZE]marketHistoryInfo
	// TODO could this be improved?!
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
