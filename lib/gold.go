package lib

// GoldPricesUpload contains the current gold prices
type GoldPricesUpload struct {
	Prices     []int `json:"Prices"`
	TimeStamps []int `json:"Timestamps"`
}
