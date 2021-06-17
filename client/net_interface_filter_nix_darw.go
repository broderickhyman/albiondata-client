// +build linux darwin

package client

import (
	"net"
)

// Gets all physical interfaces based on filter results, ignoring all VM, Loopback and Tunnel interfaces.
func getAllPhysicalInterface() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var outInterfaces []string

	for _, _interface := range interfaces {
		if _interface.Flags&net.FlagLoopback == 0 && _interface.Flags&net.FlagUp == 1 && isPhysicalInterface(_interface.HardwareAddr.String()) {
			outInterfaces = append(outInterfaces, _interface.Name)
		}
	}

	return outInterfaces, nil
}
