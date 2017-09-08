package client

import (
	"strconv"
	"strings"

	"github.com/regner/albiondata-client/lib"
	"github.com/regner/albiondata-client/log"
)

type operationReadMail struct {
	ID   int    `mapstructure:"0"`
	Body string `mapstructure:"1"`
}

func (op operationReadMail) Process(state *albionState) {
	log.Debug("Got ReadMail operation...")

	// split the mail body
	body := strings.Split(op.Body, "|")

	var notification lib.MarketNotification

	// looks like this is a market sell notification.
	switch len(body) {
	case 5:
		notification = decodeSellNotification(op, body)
	case 3:
		notification = decodeExpiryNotification(op, body)
		// this is a normal mail or something else we're not interested in
	default:
		return
	}

	upload := lib.MarketNotificationUpload{
		Type:         notification.Type(),
		Notification: notification,
	}

	log.Info("Sending a market notification to private ingest")
	sendMsgToPrivateUploaders(&upload, lib.NatsMarketNotifications, state)
}

func decodeSellNotification(op operationReadMail, body []string) lib.MarketNotification {
	notification := &lib.MarketSellNotification{}
	notification.MailID = op.ID
	notification.BuyerName = body[0]

	amount, err := strconv.Atoi(body[1])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)

		return nil
	}

	notification.Amount = amount
	notification.ItemID = body[2]

	price, err := strconv.Atoi(body[3])
	if err != nil {
		log.Error("Could not parse price in market sell notification ", err)

		return nil
	}

	notification.Price = price / 10000
	notification.TotalAfterTaxes = float32(float32(notification.Price) * float32(notification.Amount) * (1.0 - lib.SalesTax))

	return notification
}

func decodeExpiryNotification(op operationReadMail, body []string) lib.MarketNotification {
	notification := &lib.MarketExpiryNotification{}
	notification.MailID = op.ID

	amount, err := strconv.Atoi(body[0])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)

		return nil
	}

	notification.Amount = amount
	notification.ItemID = body[1]

	return notification
}
