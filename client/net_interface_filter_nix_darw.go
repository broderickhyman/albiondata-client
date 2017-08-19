// +build linux darwin

package client

import (
	"net"

	"github.com/regner/albionmarket-client/log"
)

// Gets the first physical interface based on filter results, ignoring all VM, Loopback and Tunnel interfaces
func GetFirstPhysicalInterface() string {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
		return ""
	}

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 && element.Flags&net.FlagUp == 1 && isPhysicalInterface(element.HardwareAddr.String()) {
			return element.Name
		}
	}

	return ""
}

// Gets all physical interfaces based on filter results, ignoring all VM, Loopback and Tunnel interfaces.
func GetAllPhysicalInterface() []string {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
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
