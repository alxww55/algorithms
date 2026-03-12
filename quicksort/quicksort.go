package quicksort

import (
	"slices"
)

func ExecQuickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	less := []int{}
	greater := []int{}

	for _, v := range slice[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return slices.Concat(ExecQuickSort(less), []int{pivot}, ExecQuickSort(greater))
}
