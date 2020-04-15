package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	var tests = []struct {
		Nums     []int
		Expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		actual := majorityElement(test.Nums)
		assert.Equal(t, test.Expected, actual)
	}
}
