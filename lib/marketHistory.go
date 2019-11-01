package lib

import (
	"fmt"
)

// Timescale represents the three different timescale views available for market history
type Timescale uint8

const ( // iota defaults to 0 and then incremeneted by 1
	// Hours - 0 means we are looking at the 24 hour scale
	Hours Timescale = iota
	// Days - 1 means we are looking at the 7 day scale
	Days Timescale = iota // c1 == 1
	// Weeks - 2 means we are looking at the 4 week scale
	Weeks Timescale = iota // c2 == 2
) // consts for marketHistory

func (scale Timescale) String() string {
	names := [...]string{
		"Hours",
		"Days",
		"Weeks"}

	if scale < Hours || scale > Weeks {
		return "Invalid Timescale"
	}

	return names[scale]
}

// MarketHistory contains the ItemID, Timescale, and 3 data arrays (Item amount, silver amount, and timetamp)
// These values come over the wire with indexes aligned, but are likely not sorted by time.
// Their sizes also value based on need as mentioned below.
type MarketHistory struct {
	AlbionID     uint32    `json:"AlbionID"`
	LocationID   int       `json:"LocationId"`
	QualityLevel uint8     `json:"QualityLevel"`
	Timescale    Timescale `json:"Timescale"`
	ItemAmount   uint64    `json:"ItemAmount"`
	SilverAmount uint64    `json:"SilverAmount"`
	Timestamp    uint64    `json:"Timestamp"`
	// even for the same parameter type, array type will differ depending on the size of the data values being sent.
	// For this reason, we'll be safe and use the largest expected values.
}

// StringArray for MarketHistory, duh
func (m *MarketHistory) StringArray() []string {
	return []string{
		fmt.Sprintf("%d", m.AlbionID),
		fmt.Sprintf("%d", m.LocationID),
		fmt.Sprintf("%d", m.QualityLevel),
		fmt.Sprintf("%s", m.Timescale),
		fmt.Sprintf("%d", m.ItemAmount),
		fmt.Sprintf("%d", m.SilverAmount),
		fmt.Sprintf("%d", m.Timestamp),
	}
}

type MarketHistoriesUpload struct {
	Histories []*MarketHistory `json:"MarketHistories"`
}
