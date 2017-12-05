package client

import (
	nats "github.com/nats-io/go-nats"
	"github.com/alexscott/albiondata-client/log"
)

type natsUploader struct {
	isPrivate bool
	url       string
	nc        *nats.Conn
}

// newNATSUploader creates a new NATS uploader
func newNATSUploader(url string) uploader {
	nc, _ := nats.Connect(url)

	return &natsUploader{
		url: url,
		nc:  nc,
	}
}

func (u *natsUploader) sendToIngest(body []byte, topic string) {
	if err := u.nc.Publish(topic, body); err != nil {
		log.Errorf("Error while sending ingest to nats with data: %v", err)
	}
}
