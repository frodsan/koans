package koans

import (
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	var tests = []struct {
		arr      []int32
		expected int32
	}{
		{[]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1, 2, 2, 3, 1, 2}, 5},
	}

	for _, tt := range tests {
		max := PickingNumbers(tt.arr)

		if max != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, max)
		}
	}
}

func TestPickingNumbers2(t *testing.T) {
	var tests = []struct {
		arr      []int32
		expected int32
	}{
		{[]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{[]int32{4, 6, 5, 3, 3, 1}, 3},
		{[]int32{1, 2, 2, 3, 1, 2}, 5},
	}

	for _, tt := range tests {
		max := PickingNumbers2(tt.arr)

		if max != tt.expected {
			t.Errorf("Expected %d, got %d", tt.expected, max)
		}
	}
}
