package client

import (
	"github.com/labstack/gommon/log"
	"github.com/nats-io/go-nats"
)

type natsUploader struct {
	url string
	nc  *nats.Conn
}

// newNATSUploader create a new NATS uploader :)
func newNATSUploader(url string) iuploader {
	nc, _ := nats.Connect(url)

	return &natsUploader{
		url: url,
		nc:  nc,
	}
}

func (u *natsUploader) sendToIngest(body []byte, queue string) {
	if err := u.nc.Publish(queue, body); err != nil {
		log.Errorf("Error while sending ingest to nats with data: %v", err)
	}
}
