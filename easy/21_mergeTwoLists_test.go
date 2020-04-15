package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	var tests = []struct {
		L1       *ListNode
		L2       *ListNode
		Expected []int
	}{
		{
			L1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			L2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			Expected: []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, test := range tests {
		actual := mergeTwoLists(test.L1, test.L2)
		var actualArr []int
		for actual != nil {
			actualArr = append(actualArr, actual.Val)
			actual = actual.Next
		}
		assert.Equal(t, test.Expected, actualArr)
	}
}
