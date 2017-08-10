package client

import (
	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/net"
	"golang.org/x/tools/container/intsets"
)

func findProcess(processName string) []int {
	var results []int
	var processes, _ = ps.Processes()

	for _, proc := range processes {
		if proc.Executable() == processName {
			results = append(results, proc.Pid())
		}
	}

	return results
}

func getProcessPorts(pid int) []int {
	var connections, _ = net.ConnectionsPid("all", int32(pid))
	var result = make([]int, len(connections))

	for i, c := range connections {
		result[i] = int(c.Laddr.Port)
	}

	return result

}

func diffIntSets(a []int, b []int) ([]int, []int) {
	var aSparse = intsets.Sparse{}
	var bSparse = intsets.Sparse{}

	for _, k := range a {
		aSparse.Insert(k)
	}

	for _, k := range b {
		bSparse.Insert(k)
	}

	var addedSparse = intsets.Sparse{}
	addedSparse.Difference(&bSparse, &aSparse)
	var addedSlice []int = addedSparse.AppendTo(make([]int, 0))

	var removedSparse = intsets.Sparse{}
	removedSparse.Difference(&aSparse, &bSparse)
	var removedSlice []int = removedSparse.AppendTo(make([]int, 0))

	var added = make([]int, addedSparse.Len())
	var removed = make([]int, removedSparse.Len())

	for i := 0; i < addedSparse.Len(); i++ {
		added[i] = addedSlice[i]
	}

	for i := 0; i < removedSparse.Len(); i++ {
		removed[i] = removedSlice[i]
	}

	return added, removed
}
