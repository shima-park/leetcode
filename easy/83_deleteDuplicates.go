package easy

import "math"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var v = -(math.MaxInt32 - 1)
	var prev *ListNode
	temp := head
	for temp != nil {
		// 相同时删除当前元素，将前元素Next指向后面一个元素
		if v == temp.Val {
			prev.Next = temp.Next
		} else {
			v = temp.Val
			prev = temp
		}
		temp = temp.Next
	}
	return head
}
