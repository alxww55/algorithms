package binarysearch

import (
	"fmt"
)

func ExecBinarySearch(inputSlice []int, elemToFind int) (int, error) {

	lowerBound := 0
	upperBound := len(inputSlice) - 1

	for lowerBound <= upperBound {
		middleElem := (lowerBound + upperBound) / 2
		guessedElem := inputSlice[middleElem]
		if guessedElem == elemToFind {
			return middleElem, nil
		} else if guessedElem > elemToFind {
			upperBound = middleElem - 1
		} else {
			lowerBound = middleElem + 1
		}
	}
	return -1, fmt.Errorf("Element %d not found in slice", elemToFind)
}
