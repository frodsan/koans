package koans

import "testing"

func TestArraySum(t *testing.T) {
	sum := ArraySum([]int{1, 2, 3, 4, 10, 11})

	if sum != 31 {
		t.Error("Expected 31, got ", sum)
	}
}
