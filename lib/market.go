package lib

// MarketOrder contains an order (offer or request)
type MarketOrder struct {
	ID               int    `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
	LocationID       int    `json:"LocationId"`
	QualityLevel     int    `json:"QualityLevel"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	Price            int    `json:"UnitPriceSilver"`
	Amount           int    `json:"Amount"`
	AuctionType      string `json:"AuctionType"`
	Expires          string `json:"Expires"`
}

// MarketUpload contains a list of orders and the location where the orders are from
type MarketUpload struct {
	Orders []*MarketOrder `json:"Orders"`
}
