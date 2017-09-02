package client

import (
	nats "github.com/nats-io/go-nats"
	"github.com/regner/albiondata-client/log"
)

type natsUploader struct {
	isPrivate bool
	url       string
	nc        *nats.Conn
}

// newNATSUploader creates a new NATS uploader :)
func newNATSUploader(url string, isPrivate bool) iuploader {
	nc, _ := nats.Connect(url)

	return &natsUploader{
		isPrivate: isPrivate,
		url:       url,
		nc:        nc,
	}
}

func (u *natsUploader) private() bool {
	return u.isPrivate
}

func (u *natsUploader) sendToPrivateIngest(body []byte, queue string) {
	if u.private() {
		u.sendToIngest(body, queue)
	}
}

func (u *natsUploader) sendToPublicIngest(body []byte, queue string) {
	u.sendToIngest(body, queue)
}

func (u *natsUploader) sendToIngest(body []byte, queue string) {
	if err := u.nc.Publish(queue, body); err != nil {
		log.Errorf("Error while sending ingest to nats with data: %v", err)
	}
}
