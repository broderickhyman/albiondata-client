package client

import (
	"github.com/mitchellh/mapstructure"
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
	case 67:
		operation := operationAuctionGetOffersResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 68:
		operation := operationAuctionGetRequestsResponse{}
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
