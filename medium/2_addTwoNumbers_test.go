package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		L1       []int
		L2       []int
		Expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{9}, []int{9}, []int{8, 1}},
		{[]int{1, 9}, []int{0}, []int{1, 9}},
	}

	for _, test := range tests {
		l1 := &ListNode{}
		l1.Init(test.L1...)

		l2 := &ListNode{}
		l2.Init(test.L2...)

		actual := addTwoNumbers(l1, l2).ToInts()
		assert.Equal(t, test.Expected, actual)
	}
}
