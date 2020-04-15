package easy

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := recursive(root.Left, root.Right)
	ret = append(ret, []int{root.Val})
	return ret
}

func recursive(left, right *TreeNode) [][]int {
	var arr []int
	var leftRet [][]int
	var rightRet [][]int
	var ret [][]int
	if left != nil {
		leftRet = recursive(left.Left, left.Right)
		arr = append(arr, left.Val)
	}

	if right != nil {
		rightRet = recursive(right.Left, right.Right)
		arr = append(arr, right.Val)
	}

	if len(leftRet) > 0 {
		ret = append(ret, leftRet...)
	}

	if len(rightRet) > 0 {
		ret = append(ret, rightRet...)
	}

	if len(arr) > 0 {
		ret = append(ret, arr)
	}
	return ret
}
