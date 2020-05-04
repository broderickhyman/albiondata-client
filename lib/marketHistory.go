package lib

import "fmt"

// Timescale represents the three different timescale views available for market history
type Timescale uint8

const ( // iota defaults to 0 and then incremented by 1
	// Hours - 0 means we are looking at the 24 hour scale
	Hours Timescale = 0
	// Days - 1 means we are looking at the 7 day scale
	Days Timescale = 1
	// Weeks - 2 means we are looking at the 4 week scale
	Weeks Timescale = 2
) // const for marketHistory

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

// MarketHistory contains the 3 data arrays (Item amount, silver amount, and timestamp)
// These values come over the wire with indexes aligned, but are likely not sorted by time.
// Their sizes also value based on need as mentioned below.
type MarketHistory struct {
	ItemAmount   int64 `json:"ItemAmount"`
	SilverAmount uint64 `json:"SilverAmount"`
	Timestamp    uint64 `json:"Timestamp"`
	// even for the same parameter type, array type will differ depending on the size of the data values being sent.
	// For this reason, we'll be safe and use the largest expected values.
}

// StringArray for MarketHistory
func (m *MarketHistory) StringArray() []string {
	return []string{
		fmt.Sprintf("%d", m.ItemAmount),
		fmt.Sprintf("%d", m.SilverAmount),
		fmt.Sprintf("%d", m.Timestamp),
	}
}

type MarketHistoriesUpload struct {
	AlbionId     uint32           `json:"AlbionId"`
	LocationId   int              `json:"LocationId"`
	QualityLevel uint8            `json:"QualityLevel"`
	Timescale    Timescale        `json:"Timescale"`
	Histories    []*MarketHistory `json:"MarketHistories"`
}
