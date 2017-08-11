package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/regner/albionmarket-client/client/operations"
)

func decode(params map[string]interface{}) operation {
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
	}

	return nil
}
