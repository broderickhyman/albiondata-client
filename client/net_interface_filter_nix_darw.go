// +build linux darwin

package client

import (
	"net"

	"github.com/broderickhyman/albiondata-client/log"
)

// Gets all physical interfaces based on filter results, ignoring all VM, Loopback and Tunnel interfaces.
func getAllPhysicalInterface() []string {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Panic(err)
		return nil
	}

	var outInterfaces []string

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 && element.Flags&net.FlagUp == 1 && isPhysicalInterface(element.HardwareAddr.String()) {
			outInterfaces = append(outInterfaces, element.Name)
		}
	}

	return outInterfaces
}
