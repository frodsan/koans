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

func PickingNumbers(a []int32) int32 {
	var m int32

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	for i := 0; i < len(a)-1; i++ {
		var ctr int32 = 1

		for j := i + 1; j < len(a); j++ {
			if abs(a[i]-a[j]) <= 1 {
				ctr++
			}
		}

		if m < ctr {
			m = ctr
		}
	}

	return m
}

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func PickingNumbers2(a []int32) int32 {
	arr := [100]int32{}

	for _, n := range a {
		arr[n]++
	}

	var m int32

	for i := 0; i < len(arr)-1; i++ {
		m = max(m, arr[i]+arr[i+1])
	}

	return m
}

func max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}
