package client

type operation interface {
	Process(state *albionState, uploader iuploader)
}
