package client

import (
	"strconv"
	"strings"

	"github.com/broderickhyman/albiondata-client/lib"
	"github.com/broderickhyman/albiondata-client/log"
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

	_, err := strconv.Atoi(body[3])
	if err != nil || len(body) != 4 {
		// this is a expired buy or sell mail, we're not interested in
		return
	}

	// looks like this is a market sell or buy notification.
	// need to opGetMailInfos to identify if is Buy or Sell.
	notification = decodeSellNotification(op, body)

	if notification == nil {
		return
	}

	upload := lib.MarketNotificationUpload{
		Type:         notification.Type(),
		Notification: notification,
	}

	sendMsgToPrivateUploaders(&upload, lib.NatsMarketNotifications, state)
}

func decodeSellNotification(op operationReadMail, body []string) lib.MarketNotification {
	notification := &lib.MarketSellNotification{}
	notification.MailID = op.ID
	//notification.BuyerName = body[0]

	amount, err := strconv.Atoi(body[0])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)
		return nil
	}

	notification.Amount = amount
	notification.ItemID = body[1]

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
