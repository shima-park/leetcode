package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var r = &ListNode{}
	temp := r
	for l1 != nil || l2 != nil {
		var val int
		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			val = l1.Val
			l1 = l1.Next
		} else if l2 != nil && (l1 == nil || l2.Val <= l1.Val) {
			val = l2.Val
			l2 = l2.Next
		}
		temp.Next = &ListNode{
			Val: val,
		}
		temp = temp.Next
	}

	return r.Next
}
