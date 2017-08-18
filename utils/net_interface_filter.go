package utils

import (
	"errors"
	"net"
	"strings"
)

// Compares characters until the end of lString is reached or a difference was found.
func rStrBeginsWithLstring(lString string, rString string, err *error) bool {
	if len(rString) < len(lString) {
		(*err) = errors.New("source is bigger than target string")
		return false
	}

	ls := strings.ToLower(lString)
	lt := strings.ToLower(rString)

	for i := 0; i < len(lString); i++ {
		if ls[i] != lt[i] {
			return false
		}
	}

	return true
}

// Filters the possible physical interface address by comparing it to known popular VM Software adresses
// and Teredo Tunneling Pseudo-Interface.
func isPhysicalInterface(addr *net.HardwareAddr) bool {
	var err *error = nil

	// Mac Address parts to look for, and identify non physical devices. There may be more, update me!
	var macAddrPartsToCheck []string = []string{
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

	for _, macPart := range macAddrPartsToCheck {
		if rStrBeginsWithLstring(macPart, (*addr).String(), err) {
			return false
		}
	}

	return true
}

// Gets the first physical interface based on filter results, ignoring Loopback interfaces
// Pass MAC address in XX:XX format to let it through. Otherwise leave it empty to block all VM, Loopback and Tunnel
func GetFirstPhysicalInterface(ignoreMacAddr string) (*net.Interface, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 {
			if len(ignoreMacAddr) > 0 && strings.ToLower(element.HardwareAddr.String()) == strings.ToLower(ignoreMacAddr) {
				return &element, nil
			} else if isPhysicalInterface(&element.HardwareAddr) {
				return &element, nil
			}
		}
	}

	return nil, nil
}

// Gets all physical interfaces based on filter results, ignoring Loopback interfaces.
// Pass MAC address in XX:XX format to let it through. Otherwise leave it empty to block all VM, Loopback and Tunnel
func GetAllPhysicalInterface(ignoreMacAddr string) (*[]net.Interface, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	var outInterfaces []net.Interface

	for _, element := range ifaces {
		if element.Flags&net.FlagLoopback == 0 {
			if len(ignoreMacAddr) > 0 && strings.ToLower(element.HardwareAddr.String()) == strings.ToLower(ignoreMacAddr) {
				outInterfaces = append(outInterfaces, element)
			} else if isPhysicalInterface(&element.HardwareAddr) || ignoreMacAddr == element.HardwareAddr.String() {
				outInterfaces = append(outInterfaces, element)
			}
		}
	}

	return &outInterfaces, nil
}
