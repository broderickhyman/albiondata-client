package lib

import (
	"fmt"
	"strings"
)

// GoldPricesUpload contains the current gold prices
type GoldPricesUpload struct {
	Prices     []int `json:"Prices"`
	TimeStamps []int `json:"Timestamps"`
}

func (g *GoldPricesUpload) StringArray() []string {
	var prices []string
	for _, price := range g.Prices {
		prices = append(prices, fmt.Sprintf("%d", price))
	}

	var tss []string
	for _, ts := range g.TimeStamps {
		tss = append(tss, fmt.Sprintf("%d", ts))
	}

	return []string{strings.Join(prices, ";"), strings.Join(tss, ";")}
}

func GetGoldPricesUploadJsonKeys() []string {
	return []string{
		"Prices",
		"Timestamps",
	}
}
