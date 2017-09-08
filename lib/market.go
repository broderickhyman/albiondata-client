package lib

import "fmt"

// MarketOrder contains an order (offer or request)
type MarketOrder struct {
	ID               int    `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
	GroupTypeId      string `json:"ItemGroupTypeId"`
	LocationID       int    `json:"LocationId"`
	QualityLevel     int    `json:"QualityLevel"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	Price            int    `json:"UnitPriceSilver"`
	Amount           int    `json:"Amount"`
	AuctionType      string `json:"AuctionType"`
	Expires          string `json:"Expires"`
}

func (m *MarketOrder) StringArray() []string {
	return []string{
		fmt.Sprintf("%d", m.ID),
		m.ItemID,
		fmt.Sprintf("%d", m.LocationID),
		fmt.Sprintf("%d", m.QualityLevel),
		fmt.Sprintf("%d", m.EnchantmentLevel),
		fmt.Sprintf("%d", m.Price),
		fmt.Sprintf("%d", m.Amount),
		m.AuctionType,
		m.Expires,
	}
}

const (
	SalesTax = 0.02
)

type MarketNotificationType string

const (
	SalesNotification  MarketNotificationType = "SalesNotification"
	ExpiryNotification                        = "ExpiryNotification"
)

type MarketNotification interface {
	Type() MarketNotificationType
}

type MarketSellNotification struct {
	MailID          int     `json:"Id"`
	BuyerName       string  `json:"BuyerName"`
	ItemID          string  `json:"ItemTypeId"`
	Amount          int     `json:"Amount"`
	Price           int     `json:"UnitPriceSilver"`
	TotalAfterTaxes float32 `json:"TotalAfterTaxes"`
}

type MarketExpiryNotification struct {
	MailID int    `json:"Id"`
	ItemID string `json:"ItemTypeId"`
	Amount int    `json:"Amount"`
}

func (m *MarketSellNotification) Type() MarketNotificationType {
	return SalesNotification
}

func (m *MarketExpiryNotification) Type() MarketNotificationType {
	return ExpiryNotification
}

type MarketNotificationUpload struct {
	PrivateUpload
	Type         MarketNotificationType `json: "NotificationType"`
	Notification MarketNotification     `json:"Notification"`
}

// MarketUpload contains a list of orders
type MarketUpload struct {
	Orders []*MarketOrder `json:"Orders"`
}
