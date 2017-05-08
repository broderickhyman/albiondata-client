package shared

import "time"

type InjestPostRequest struct {
	MarketItems []string `json:"marketItems"`
}

type MarketUpdate struct {
	MarketItems []MarketItem `json:"MarketItems"`
	IngestTime  time.Time    `json:"IngestTime"`
}

type MarketItem struct {
	ID               int    `json:"Id"`
	UnitPrice        int    `json:"UnitPriceSilver"`
	TotalPrice       int    `json:"TotalPriceSilver"`
	Amount           int    `json:"Amount"`
	Tier             int    `json:"Teir"`
	ItemTypeID       string `json:"ItemTypeId"`
	ItemGroupTypeID  string `json:"ItemGroupTypeId"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	QualityLevel     int    `json:"QualityLevel"`
	Expires          string `json:"Expires"`
}
