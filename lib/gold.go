package lib

// GoldInfoUpload contains the current gold prices
type GoldInfoUpload struct {
	Prices     []int `json:"prices"`
	TimeStamps []int `json:"timestamps"`
}
