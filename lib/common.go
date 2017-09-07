package lib

type PrivateUpload struct {
	CharacterId   string `json:"CharacterId"`
	CharacterName string `json:"CharacterName"`
}

func (p *PrivateUpload) Personalize(id string, name string) {
	p.CharacterId = id
	p.CharacterName = name
}

type PersonalizedUpload interface {
	Personalize(string, string)
}
