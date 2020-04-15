package easy

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		Nums     []int
		Target   int
		Expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range tests {
		actual := twoSum(test.Nums, test.Target)
		sort.Ints(actual)
		assert.Equal(t, test.Expected, actual)
	}
}
