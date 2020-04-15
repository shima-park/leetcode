package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		Nums1    []int
		Nums2    []int
		Expected float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		actual := findMedianSortedArrays(test.Nums1, test.Nums2)
		assert.Equal(t, test.Expected, actual)
	}
}
