package lib

// MapDataUpload contains information on zone maps
type MapDataUpload struct {
	ZoneID          int      `json:"ZoneID"`
	BuildingType    []int    `json:"BuildingType"`
	AvailableFood   []int    `json:"AvailableFood"`
	Reward          []int    `json:"Reward"`
	AvailableSilver []int    `json:"AvailableSilver"`
	Owners          []string `json:"Owners"`
	Buildable       []bool   `json:"Buildable"`
	IsForSale       []bool   `json:"IsForSale"`
	BuyPrice        []int    `json:"BuyPrice"`
}
