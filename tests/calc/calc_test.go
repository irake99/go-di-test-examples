// calc_test.go
package calc

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 1, -1, 1}, 0},
	}

	for _, test := range tests {
		result := Sum(test.numbers)
		if result != test.expected {
			t.Errorf("For %v, expected %d, but got %d", test.numbers, test.expected, result)
		}
	}
}
