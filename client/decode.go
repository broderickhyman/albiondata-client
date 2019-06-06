package client

import (
	"albiondata-client/lib"
	"albiondata-client/log"
	"encoding/hex"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func decodeRequest(params map[string]interface{}) (operation operation, err error) {
	if _, ok := params["253"]; !ok {
		return nil, nil
	}

	code := params["253"].(int16)

	switch OperationType(code) {
	case opGetGameServerByCluster:
		operation = &operationGetGameServerByCluster{}
	case opAuctionGetOffers:
		operation = &operationAuctionGetOffers{}
	case opGetClusterMapInfo:
		operation = &operationGetClusterMapInfo{}
	case opGoldMarketGetAverageInfo:
		operation = &operationGoldMarketGetAverageInfo{}
	case opRealEstateGetAuctionData:
		operation = &operationRealEstateGetAuctionData{}
	case opRealEstateBidOnAuction:
		operation = &operationRealEstateBidOnAuction{}
	default:
		return nil, nil
	}

	err = decodeParams(params, operation)

	return operation, err
}

func decodeResponse(params map[string]interface{}) (operation operation, err error) {
	if _, ok := params["253"]; !ok {
		return nil, nil
	}

	code := params["253"].(int16)

	switch OperationType(code) {
	case opJoin:
		operation = &operationJoinResponse{}
	case opAuctionGetOffers:
		operation = &operationAuctionGetOffersResponse{}
	case opAuctionGetRequests:
		operation = &operationAuctionGetRequestsResponse{}
	case opReadMail:
		operation = &operationReadMail{}
	case opGetClusterMapInfo:
		operation = &operationGetClusterMapInfoResponse{}
	case opGoldMarketGetAverageInfo:
		operation = &operationGoldMarketGetAverageInfoResponse{}
	case opRealEstateGetAuctionData:
		operation = &operationRealEstateGetAuctionDataResponse{}
	case opRealEstateBidOnAuction:
		operation = &operationRealEstateBidOnAuctionResponse{}
	default:
		return nil, nil
	}

	err = decodeParams(params, operation)

	return operation, err
}

func decodeEvent(params map[string]interface{}) (event operation, err error) {
	if _, ok := params["252"]; !ok {
		return nil, nil
	}

	eventType := params["252"].(int16)

	switch eventType {
	// case evRespawn: //TODO: confirm this eventCode (old 77)
	// 	event = &eventPlayerOnlineStatus{}
	// case evCharacterStats: //TODO: confirm this eventCode (old 114)
	// 	event = &eventSkillData{}
	default:
		return nil, nil
	}

	err = decodeParams(params, event)

	return event, err
}

func decodeParams(params interface{}, operation operation) error {
	convertGameObjects := func(from reflect.Type, to reflect.Type, v interface{}) (interface{}, error) {
		if from == reflect.TypeOf([]int8{}) && to == reflect.TypeOf(lib.CharacterID("")) {
			log.Debug("Parsing character ID from mixed-endian UUID")

			return decodeCharacterID(v.([]int8)), nil
		}

		return v, nil
	}

	config := mapstructure.DecoderConfig{
		DecodeHook: convertGameObjects,
		Result:     operation,
	}

	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}

	err = decoder.Decode(params)

	return err
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
