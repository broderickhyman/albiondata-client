package client

import (
	"github.com/regner/albiondata-client/log"
)

type iuploader interface {
	sendToIngest(body []byte, queue string)
}

type noopUploader struct {
}

func (u *noopUploader) sendToIngest(body []byte, queue string) {
	log.Debugf("Got a noop request to queue: %s", queue)
	return
}
