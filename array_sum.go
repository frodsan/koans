package koans

func ArraySum(arr []int) int {
	var sum int

	for _, n := range arr {
		sum += n
	}

	return sum
}
