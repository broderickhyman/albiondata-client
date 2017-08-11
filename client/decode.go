package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/regner/albionmarket-client/client/operations"
)

func decodeRequest(params map[string]interface{}) operation {
	if _, ok := params["253"]; !ok {
		return nil
	}

	code := params["253"].(int16)

	switch code {
	case 10:
		operation := operations.GetGameServerByCluster{}
		mapstructure.Decode(params, &operation)

		return operation
	case 67:
		operation := operations.AuctionGetOffers{}
		mapstructure.Decode(params, &operation)

		return operation
	case 217:
		operation := operations.GoldMarketGetAverageInfo{}
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
		operation := operations.AuctionGetOffersResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 68:
		operation := operations.AuctionGetRequestsResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	case 217:
		operation := operations.GoldMarketGetAverageInfoResponse{}
		mapstructure.Decode(params, &operation)

		return operation
	}

	return nil
}
