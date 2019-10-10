package lib

type PrivateUpload struct {
	CharacterId   CharacterID `json:"CharacterId"`
	CharacterName string      `json:"CharacterName"`
}

func (p *PrivateUpload) Personalize(id CharacterID, name string) {
	p.CharacterId = id
	p.CharacterName = name
}

type PersonalizedUpload interface {
	Personalize(CharacterID, string)
}

// Represents a character identifier in its UUID-style string format
type CharacterID string
