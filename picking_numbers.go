package koans

import (
	"sort"
)

/*

Source: Hacker Rank (https://www.hackerrank.com/challenges/picking-numbers/)

Given an array of integers, find and print the maximum number of integers you can select
from the array such that the absolute difference between any two of the chosen integers
is less than or equal to `1`. For example, if your array is `[1, 1, 2, 2, 4, 4, 5, 5, 5],
you can create two subarrays meeting the criterion: [1, 1, 2, 2] and [4, 4, 5, 5, 5].
The maximum length subarray has `5` elements.

*/

func pickNumbers(arr []int) int {
	max := 0

	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		ctr := 1

		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) <= 1 {
				ctr++
			}
		}

		if max < ctr {
			max = ctr
		}
	}

	return max
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
