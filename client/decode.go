package client

import (
	"encoding/hex"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

func decodeRequest(params map[string]interface{}) operation {
	if _, ok := params["253"]; !ok {
		return nil
	}

	code := params["253"].(int16)

	switch code {
	case 10:
		operation := operationGetGameServerByCluster{}
		mapstructure.Decode(params, &operation)

		return operation
	case 67:
		operation := operationAuctionGetOffers{}
		mapstructure.Decode(params, &operation)

		return operation
	case 166:
		operation := operationGetClusterMapInfo{}
		mapstructure.Decode(params, &operation)

		return operation
	case 217:
		operation := operationGoldMarketGetAverageInfo{}
		mapstructure.Decode(params, &operation)

		return operation
	case 232:
		operation := operationRealEstateGetAuctionData{}
		mapstructure.Decode(params, &operation)

		return operation
	case 233:
		operation := operationRealEstateBidOnAuction{}
		mapstructure.Decode(params, &operation)

		return operation
	}

	return nil
}

func decodeResponse(params map[string]interface{}) operation {
	if _, ok := params["253"]; !ok {
		return nil
	}

	code := params["253"].(int16)

	switch code {
	case 2:
		operation := operationJoinResponse{}
		decodeParams(params, &operation)
		return operation
	case 67:
		operation := operationAuctionGetOffersResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 68:
		operation := operationAuctionGetRequestsResponse{}
		mapstructure.Decode(params, &operation)

		return operation

	case 147:
		operation := operationReadMail{}
		mapstructure.Decode(params, &operation)
		return operation

	case 166:
		operation := operationGetClusterMapInfoResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 217:
		operation := operationGoldMarketGetAverageInfoResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 232:
		operation := operationRealEstateGetAuctionDataResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 233:
		operation := operationRealEstateBidOnAuctionResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	}

	return nil
}

func decodeEvent(params map[string]interface{}) operation {
	if _, ok := params["252"]; !ok {
		return nil
	}

	eventType := params["252"].(int16)

	switch eventType {
	case 77:
		event := eventPlayerOnlineStatus{}
		err := decodeParams(params, &event)
		log.Debug(err)

		return event
	case 114:
		event := eventSkillData{}
		mapstructure.Decode(params, &event)

		return event
	}

	return nil
}

func decodeParams(params interface{}, out interface{}) error {
	convertGameObjects := func(from reflect.Type, to reflect.Type, v interface{}) (interface{}, error) {
		if from == reflect.TypeOf([]int8{}) && to == reflect.TypeOf(lib.CharacterID("")) {
			log.Debug("Parsing character ID from mixed-endian UUID")

			return decodeCharacterID(v.([]int8)), nil
		}

		return v, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: convertGameObjects,
		Result:     out,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}

	return decoder.Decode(params)
}

func decodeCharacterID(array []int8) lib.CharacterID {
	/* So this is a UUID, which is stored in a 'mixed-endian' format.
	The first three components are stored in little-endian, the rest in big-endian.
	See https://en.wikipedia.org/wiki/Universally_unique_identifier#Encoding.
	By default, our int array is read as big-endian, so we need to swap the first
	three components of the UUID
	*/
	b := make([]byte, len(array))

	// First, convert to byte
	for k, v := range array {
		b[k] = byte(v)
	}

	// swap first component
	b[0], b[1], b[2], b[3] = b[3], b[2], b[1], b[0]

	// swap second component
	b[4], b[5] = b[5], b[4]

	// swap third component
	b[6], b[7] = b[7], b[6]

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

	return lib.CharacterID(buf[:])
}
