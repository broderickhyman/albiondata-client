package utils

import (
	"github.com/google/gopacket"
)

func IsChannelSignaled(quit chan bool) bool {
	select {
			case <-quit:
				return true
			default:
				return false
		}
}

func PollChannel(channel chan gopacket.Packet) gopacket.Packet {
	select {
			case v := <- channel:
				return v
			default:
				return nil
		}
}
