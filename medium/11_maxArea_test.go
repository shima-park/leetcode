package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		Height   []int
		Expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 2}, 1},
	}

	for _, test := range tests {
		actual := maxArea(test.Height)
		assert.Equal(t, test.Expected, actual,
			"Test Case: %v", test.Height,
		)
	}
}
