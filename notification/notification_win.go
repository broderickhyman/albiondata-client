// +build windows

package notification

import (
	"github.com/alexscott/albiondata-client/log"
	toast "gopkg.in/toast.v1"
)

func Push(msg string) {
	note := toast.Notification{AppID: "Albion Data Client", Title: "Albion Data Client", Message: msg}

	err := note.Push()
	if err != nil {
		log.Error(err)
	}
}
