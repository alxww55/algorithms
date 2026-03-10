package recursion

func RecursiveSum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return slice[0] + RecursiveSum(slice[1:])
}
