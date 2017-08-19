package client

import (
	nats "github.com/nats-io/go-nats"
	"github.com/regner/albiondata-client/log"
)

type natsUploader struct {
	url string
	nc  *nats.Conn
}

// newNATSUploader creates a new NATS uploader :)
func newNATSUploader(url string) iuploader {
	nc, _ := nats.Connect(url)

	return &natsUploader{
		url: url,
		nc:  nc,
	}
}

func (u *natsUploader) sendToIngest(body []byte, queue string) {
	rawQueue := queue + ".raw"

	if err := u.nc.Publish(rawQueue, body); err != nil {
		log.Errorf("Error while sending ingest to nats with data: %v", err)
	}
}
