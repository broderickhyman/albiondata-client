package lib

import (
	"encoding/base64"
	"strings"

	"github.com/google/uuid"
)

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

func (c *CharacterID) Base64() (string, error) {
	// since it is now properly formatted, we can use the UUID packet
	UUID, err := uuid.Parse(string(*c))
	if err != nil {
		return "", err
	}

	b, err := UUID.MarshalBinary()
	if err != nil {
		return "", err
	}

	// calculate base64
	b64 := base64.RawStdEncoding.EncodeToString(b)

	// additionally, it seems that '+' are replaced by '-' and '/' with '_'
	b64 = strings.Replace(b64, "+", "-", -1)
	b64 = strings.Replace(b64, "/", "_", -1)

	return b64, nil
}
