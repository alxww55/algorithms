package selectionsort

import (
	"slices"
)

func findSmallest(slice []int) int {
	smallestElem := slice[0]
	smallestIndex := 0

	for i, v := range slice {
		if v < smallestElem {
			smallestElem = v
			smallestIndex = i
		}
	}
	return smallestIndex
}

func ExecSelectionSort(inputSlice []int) []int {
	outputSlice := []int{}

	for len(inputSlice) > 0 {
		smallestElem := findSmallest(inputSlice)
		outputSlice = append(outputSlice, inputSlice[smallestElem])
		inputSlice = slices.Delete(inputSlice, smallestElem, smallestElem+1)
	}

	return outputSlice
}
