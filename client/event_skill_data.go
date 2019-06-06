package client

import (
	"albiondata-client/lib"
	"albiondata-client/log"
	"strconv"
)

type eventSkillData struct {
	SkillIds    []int     `mapstructure:"1"`
	Levels      []int     `mapstructure:"2"`
	Percentages []float64 `mapstructure:"3"`
	Fame        []string  `mapstructure:"4"`
}

func (event eventSkillData) Process(state *albionState) {
	log.Debug("Got skill data event...")

	skills := []*lib.Skill{}

	for k, _ := range event.SkillIds {
		skill := &lib.Skill{}
		skill.ID = event.SkillIds[k]
		skill.Level = event.Levels[k]
		skill.PercentNextLevel = event.Percentages[k]
		// for some reason, the value is enclosed in [[]]. trying to get rid of them
		fame, err := strconv.Atoi(event.Fame[k][2 : len(event.Fame[k])-2])
		if err != nil {
			log.Error("Could not parse fame value. ", err)
			continue
		}
		skill.Fame = fame

		skills = append(skills, skill)
	}

	if len(skills) < 1 {
		return
	}

	upload := lib.SkillsUpload{
		Skills: skills,
	}

	log.Infof("Sending %d skills of %v to ingest", len(skills), state.CharacterName)
	sendMsgToPrivateUploaders(&upload, lib.NatsSkillData, state)
}
