package client

import (
	"github.com/mitchellh/mapstructure"
)

func decode(params map[string]interface{}) operation {
	if _, ok := params["253"]; !ok {
		return nil
	}

	code := params["253"].(int16)

	switch code {
	case 67:
		operation := requestBuyOrders{}
		mapstructure.Decode(params, &operation)

		return operation
	default:
		return nil
	}
}
