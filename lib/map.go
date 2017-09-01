package lib

import "fmt"

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

func (m *MapDataUpload) StringArrays() [][]string {
	result := [][]string{}
	for i := range m.BuildingType {
		isForSale := false
		if len(m.IsForSale) > i {
			isForSale = m.IsForSale[i]
		}

		buyPrice := 0
		if len(m.BuyPrice) > i {
			buyPrice = m.BuyPrice[i]
		}

		result = append(result, []string{
			fmt.Sprintf("%d", m.ZoneID),
			fmt.Sprintf("%d", m.BuildingType[i]),
			fmt.Sprintf("%d", m.AvailableFood[i]),
			fmt.Sprintf("%d", m.Reward[i]),
			fmt.Sprintf("%d", m.AvailableSilver[i]),
			m.Owners[i],
			fmt.Sprintf("%t", m.Buildable[i]),
			fmt.Sprintf("%t", isForSale),
			fmt.Sprintf("%d", buyPrice),
		})
	}

	return result
}
