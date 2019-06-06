// +build windows

package notification

import (
	"albiondata-client/log"

	"gopkg.in/toast.v1"
)

func Push(msg string) {
	note := toast.Notification{AppID: "Albion Data Client", Title: "Albion Data Client", Message: msg}

	err := note.Push()
	if err != nil {
		log.Error(err)
	}
}
