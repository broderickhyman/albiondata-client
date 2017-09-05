package client

type uploader interface {
	sendToIngest(body []byte, topic string)
}
