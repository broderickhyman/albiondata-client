package operations

import (
	"github.com/hmadison/ao_spectator/photon"
	"github.com/mitchellh/mapstructure"
)

func Decode(params map[string]interface {}) interface{} {
	if _, ok := params["253"]; !ok {
		return nil
	}

	var code = params["253"].(int16)

	switch code {

		case 67:
			var operation = RequestBuyOrders{}
			mapstructure.Decode(params, &operation)
			return operation

		default:
			return nil
	}
}
