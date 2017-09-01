package lib

import (
	"fmt"
	"time"
)

// GoldPricesUpload contains the current gold prices
type GoldPricesUpload struct {
	Prices     []int   `json:"Prices"`
	TimeStamps []int64 `json:"Timestamps"`
}

func (g *GoldPricesUpload) StringArrays() [][]string {
	result := [][]string{}

	for i := range g.Prices {
		result = append(result, []string{
			fmt.Sprintf("%d", g.Prices[i]),
			time.Unix(g.TimeStamps[i], 0).Format(time.RFC3339),
		})
	}

	return result
}
