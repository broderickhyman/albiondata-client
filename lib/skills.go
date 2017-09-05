package lib

// Skill contains a skill
type Skill struct {
	ID               int     `json:"Id"`
	Level            int     `json:"Level"`
	PercentNextLevel float64 `json:"PercentNextLevel"`
	Fame             int     `json:"Fame"`
}

// SkillsUpload contains a list of skills
type SkillsUpload struct {
	CharacterId   string   `json:"CharacterId"`
	CharacterName string   `json:"CharacterName"`
	Skills        []*Skill `json:"Skills"`
}
