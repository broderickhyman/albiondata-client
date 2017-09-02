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
		log.Debugf("Got a noop request to private queue: %s", queue)
	}
}

func (u *noopUploader) sendToPublicIngest(body []byte, queue string) {
	log.Debugf("Got a noop request to public queue: %s", queue)
}
