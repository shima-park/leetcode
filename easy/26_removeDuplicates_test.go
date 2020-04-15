package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	var tests = []struct {
		Nums     []int
		Expected int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range tests {
		actual := removeDuplicates(test.Nums)
		assert.Equal(t, test.Expected, actual)
	}

}
