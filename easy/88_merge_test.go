package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	arr := make([]int, 2)
	arr[0] = 2
	var tests = []struct {
		Nums1    []int
		M        int
		Nums2    []int
		N        int
		Expected []int
	}{
		{make([]int, 1), 0, []int{1}, 1, []int{1}},
		{arr, 1, []int{1}, 1, []int{1, 2}},
		{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		merge(test.Nums1, test.M, test.Nums2, test.N)
		assert.Equal(t, test.Expected, test.Nums1)
	}
}
