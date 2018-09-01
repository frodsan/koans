package koans

import (
	"testing"
)

func TestPickNumbers(t *testing.T) {
	var tests = []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{[]int{4, 6, 5, 3, 3, 1}, 3},
		{[]int{1, 2, 2, 3, 1, 2}, 5},
	}

	for _, tt := range tests {
		max := pickNumbers(tt.arr)

		if max != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, max)
		}
	}
}