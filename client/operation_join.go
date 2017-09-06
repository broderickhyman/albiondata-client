package client

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/regner/albiondata-client/log"
)

type operationJoinResponse struct {
	CharacterIDRaw     []int  `mapstructure:"1"`
	GuildIDRaw         []int  `mapstructure:"45"`
	CharacterName      string `mapstructure:"2"`
	CharacterPartsJSON string `mapstructure:"6"`
	Location           string `mapstructure:"7"`
	Edition            string `mapstructure:"38"`
	GuildName          string `mapstructure:"47"`
}

func (op operationJoinResponse) Process(state *albionState) {
	log.Debugf("Got JoinResponse operation...")

	loc, err := strconv.Atoi(op.Location)
	if err != nil {
		log.Debugf("Unable to convert zoneID to int. Probably an instance.. ZoneID: %v", op.Location)
		return
	}

	state.LocationId = loc
	log.Debugf("Updating player location to %v.", loc)

	characterIdBase64, characterId := calculateId(op.CharacterIDRaw)

	state.CharacterIdBase64 = characterIdBase64
	state.CharacterId = characterId
	log.Debugf("Updating player ID to %v (%v).", characterIdBase64, characterId)

	state.CharacterName = op.CharacterName
	log.Debugf("Updating player to %v.", op.CharacterName)
}

func calculateId(array []int) (string, string) {
	/* So this is a UUID, which is stored in a 'mixed-endian' format.
	The first three components are stored in little-endian, the rest in big-endian.
	See https://en.wikipedia.org/wiki/Universally_unique_identifier#Encoding.
	By default, our int array is read as big-endian, so we need to swap the first
	three components of the UUID
	*/
	b := make([]byte, len(array))

	// First, convert to byte
	for k, v := range array {
		b[k] = byte(v & 0xff)
	}

	// swap first component
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]

	// swap second component
	b[4], b[5] = b[5], b[4]

	// swap third component
	b[6], b[7] = b[7], b[6]

	// calculate base64
	b64 := base64.RawStdEncoding.EncodeToString(b)

	// additionally, it seems that '+' are replaced by '-' and '/' with '_'
	b64 = strings.Replace(b64, "+", "-", -1)
	b64 = strings.Replace(b64, "/", "_", -1)

	// format it UUID-style
	var buf [36]byte
	hex.Encode(buf[:], b[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], b[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], b[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], b[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], b[10:])

	return b64, string(buf[:])
}
