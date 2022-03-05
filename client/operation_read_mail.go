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
	var notification lib.MarketNotification

	// split the mail body
	body := strings.Split(op.Body, "|")

	mailInfo := MailInfos.getMailInfo(op.ID)
	if mailInfo == nil {
		log.Info("Mail Info is not valid. Please transition zones or click at notification, so the mails can be loaded.")
		return
	}

	if mailInfo.OrderType == "MARKETPLACE_SELLORDER_FINISHED_SUMMARY" {
		log.Debug("Read finished sell order.")
		notification = decodeSellNotification(op, body)
	} else if mailInfo.OrderType == "MARKETPLACE_SELLORDER_EXPIRED_SUMMARY" {
		log.Debug("Read expired sell order.")
		notification = decodeExpiryNotification(op, body)
	}

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

	sold, err := strconv.Atoi(body[0])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)
		return nil
	}

	amount, err := strconv.Atoi(body[1])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)
		return nil
	}

	price, err := strconv.Atoi(body[2])
	if err != nil {
		log.Error("Could not parse price in market sell notification ", err)
		return nil
	}

	notification.Amount = amount
	notification.ItemID = body[1]
	notification.Price = price / 10000
	notification.Sold = sold

	return notification
}
