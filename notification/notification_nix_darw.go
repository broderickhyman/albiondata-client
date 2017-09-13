// +build linux darwin

package notification

import "github.com/ctcpip/notifize"

func Push(msg string) {
	notifize.Display("Albion Data Client", msg, false, "")
}
