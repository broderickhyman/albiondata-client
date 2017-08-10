package client

import (
	"github.com/regner/albionmarket-client/client/albionstate"
)

type operation interface {
	Process(state *albionstate.AlbionState)
}
