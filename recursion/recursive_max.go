package recursion

func RecursiveMax(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	if len(slice) == 1 {
		return slice[0]
	}
	return max(slice[0], RecursiveMax(slice[1:]))
}
