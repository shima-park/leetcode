package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left != nil && root.Right != nil {
		return equalsTree(root.Left, root.Right)
	}

	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}

	return true
}

func equalsTree(left *TreeNode, right *TreeNode) bool {
	if (left != nil && right == nil) || (left == nil && right != nil) {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	if left.Val == right.Val {
		if !equalsTree(left.Left, right.Right) {
			return false
		}

		if !equalsTree(left.Right, right.Left) {
			return false
		}
		return true
	}
	return false
}
