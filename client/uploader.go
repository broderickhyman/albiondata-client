package client

import (
	"github.com/regner/albiondata-client/log"
)

type iuploader interface {
	private() bool
	sendToPublicIngest(body []byte, queue string)
	sendToPrivateIngest(body []byte, queue string)
}

func newNOOPUploader(isPrivate bool) iuploader {
	return &noopUploader{
		isPrivate: isPrivate,
	}
}

type noopUploader struct {
	isPrivate bool
}

func (u *noopUploader) private() bool {
	return u.isPrivate
}

func (u *noopUploader) sendToPrivateIngest(body []byte, queue string) {
	if u.private() {
		u.sendToIngest(body, queue, "PRIVATE")
	}
}

func (u *noopUploader) sendToPublicIngest(body []byte, queue string) {
	if u.private() {
		u.sendToIngest(body, queue, "PRIVATE")
	} else {
		u.sendToIngest(body, queue, "PUBLIC")
	}
}

func (u *noopUploader) sendToIngest(body []byte, queue string, privOrPub string) {
	log.Debugf("Got a noop request to %s queue: %s", privOrPub, queue)
}
