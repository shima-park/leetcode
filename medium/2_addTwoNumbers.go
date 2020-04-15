package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Init(arr ...int) {
	n := l
	for i, v := range arr {
		if i == 0 {
			n.Val = v
			continue
		}
		n.Next = &ListNode{
			Val: v,
		}
		n = n.Next
	}
}

func (l *ListNode) Print() {
	r := l
	for {
		if r == nil {
			fmt.Println()
			return
		}
		fmt.Print(r.Val)
		r = r.Next
	}
}

func (l *ListNode) ToInts() []int {
	r := l
	var ints []int
	for {
		if r == nil {
			return ints
		}
		ints = append(ints, r.Val)
		r = r.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	carry := 0
	tempNode := root

	getVal := func(l *ListNode) int {
		if l != nil {
			return l.Val
		}
		return 0
	}

	getNext := func(l *ListNode) *ListNode {
		if l != nil {
			return l.Next
		}
		return nil
	}

	for l1 != nil || l2 != nil {
		sum := carry + getVal(l1) + getVal(l2)

		carry = sum / 10
		tempNode.Next = &ListNode{
			Val: sum % 10,
		}
		tempNode = tempNode.Next
		l1, l2 = getNext(l1), getNext(l2)
	}

	if carry > 0 {
		tempNode.Next = &ListNode{
			Val: carry,
		}
	}
	return root.Next
}
