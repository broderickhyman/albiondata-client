package utils

import (
	"golang.org/x/tools/container/intsets"
)

func DiffIntSets(a []int, b []int) ([]int, []int) {
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
