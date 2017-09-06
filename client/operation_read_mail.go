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

	// split the mailbody
	array := strings.Split(op.Body, "|")

	// looks like this is not a market sell notification. for now we're only interested in those
	if len(array) != 5 {
		return
	}

	notification := &lib.MarketSellNotification{}
	notification.MailID = op.ID
	notification.BuyerName = array[0]

	amount, err := strconv.Atoi(array[1])
	if err != nil {
		log.Error("Could not parse amount in market sell notification ", err)
		return
	}

	notification.Amount = amount
	notification.ItemID = array[2]

	price, err := strconv.Atoi(array[3])
	if err != nil {
		log.Error("Could not parse price in market sell notification ", err)
		return
	}

	notification.Price = price / 10000
	notification.TotalAfterTaxes = float32(float32(notification.Price) * float32(notification.Amount) * (1.0 - lib.SalesTax))

	upload := lib.MarketNotificationUpload{
		Type:         notification.Type(),
		Notification: notification,
	}

	log.Info("Sending a market notification to private ingest")
	sendMsgToPrivateUploaders(&upload, lib.NatsMarketNotifications, state)
}
