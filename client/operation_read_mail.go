package client

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/albiondata-client/lib"
	"github.com/albiondata-client/log"
)

type operationReadMail struct {
	ID   int    `mapstructure:"0"`
	Body string `mapstructure:"1"`
}

func (op operationReadMail) Process(state *albionState) {

	// split the mail body
	body := strings.Split(op.Body, "|")

	var notification lib.MarketNotification
	log.Infof("Got ReadMail operation %v", body)
	// looks like this is a market sell notification.
	switch len(body) {
	case 5:
		notification = decodeSellNotification(op, body, state)
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

	sendMsgToPrivateUploaders(&upload, lib.NatsMarketNotifications, state)
}

func decodeSellNotification(op operationReadMail, body []string, state *albionState) lib.MarketNotification {

	log.Infof("trying to decode stuff")
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

	sendMsgToCSV(notification, lib.NatsMarketNotifications, state)

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

func sendMsgToCSV(upload *lib.MarketSellNotification, topic string, state *albionState) {
	log.Info("Sending to Csv")
	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	//USE CORRECT FILE MODE ^^^^^^^^^^^^^^^^^^^
	defer file.Close()

	writer := csv.NewWriter(file)

	var row []string
	row = append(row, upload.ItemID)
	row = append(row, strconv.Itoa(upload.Amount))
	row = append(row, strconv.Itoa(upload.Price))
	row = append(row, strconv.Itoa(int(upload.TotalAfterTaxes)))

	writer.Write(row)
	writer.Flush()

	if err != nil {
		log.Error("unable to create file")
	}
	return
}
