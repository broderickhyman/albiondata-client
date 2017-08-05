package utils

import (
	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/net"
)

func FindProcess(processName string) []int {
	var results []int
	var processes, _ = ps.Processes()

	for _, proc := range processes {
		if proc.Executable() == processName {
			results = append(results, proc.Pid())
		}
	}

	return results
}

func GetProcessPorts(pid int) []int {
	var connections, _ = net.ConnectionsPid("any", int32(pid))
	var result = make([]int, len(connections))

	for i, c := range connections {
		result[i] = int(c.Laddr.Port)
	}

	return result

}
