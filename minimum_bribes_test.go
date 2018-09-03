package koans

import (
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	var tests = []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 1, 5, 3, 4}, 3},
		//{[]int{1, 2, 5, 3, 7, 8, 6, 4}, 7},
	}

	for _, tt := range tests {
		max, _ := MinimumBribes(tt.arr)

		if max != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, max)
		}
	}
}

// func TestMinimumBribesTooChaotic(t *testing.T) {
// 	var tests = []struct {
// 		arr []int
// 	}{
// 		{[]int{2, 5, 1, 3, 4}},
// 		{[]int{5, 1, 2, 3, 7, 8, 6, 4}},
// 	}

// 	for _, tt := range tests {
// 		_, err := MinimumBribes(tt.arr)

// 		if err == nil {
// 			t.Errorf("Expected `Too chaotic` error")
// 		}
// 	}
// }
