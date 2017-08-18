package client

import (
	"net"
	"strings"
	
	"github.com/regner/albionmarket-client/log"
)

// Mac Address parts to look for, and identify non physical devices. There may be more, update me!
var macAddrPartsToFilter []string = []string{
	"00:03:FF",             // Microsoft Hyper-V, Virtual Server, Virtual PC
	"0A:00:27",             // VirtualBox
	"00:00:00:00:00",       // Teredo Tunneling Pseudo-Interface
	"00:50:56", "00:1C:14", // VMware ESX 3, Server, Workstation, Player
	"00:0C:29", "00:05:69", // VMware ESX 3, Server, Workstation, Player
	"00:1C:42", // Microsoft Hyper-V, Virtual Server, Virtual PC
	"00:0F:4B", // Virtual Iron 4
	"00:16:3E", // Red Hat Xen, Oracle VM, XenSource, Novell Xen
	"08:00:27", // Sun xVM VirtualBox
}

// Filters the possible physical interface address by comparing it to known popular VM Software adresses
// and Teredo Tunneling Pseudo-Interface.
func isPhysicalInterface(addr *net.HardwareAddr) bool {
	for _, macPart := range macAddrPartsToFilter {
		if strings.HasPrefix(strings.ToLower((*addr).String()), strings.ToLower(macPart)) {
			return false
		}
	}

	return true
}

// Gets the first physical interface based on filter results, ignoring all VM, Loopback and Tunnel interfaces
func GetFirstPhysicalInterface() *net.Interface {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 && isPhysicalInterface(&element.HardwareAddr) {
			return &element
		}
	}

	return nil
}

// Gets all physical interfaces based on filter results, ignoring all VM, Loopback and Tunnel interfaces.
func GetAllPhysicalInterface() *[]net.Interface {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	var outInterfaces []net.Interface

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 && isPhysicalInterface(&element.HardwareAddr) {
			outInterfaces = append(outInterfaces, element)
		}
	}

	return &outInterfaces
}
