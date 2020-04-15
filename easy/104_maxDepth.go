package easy

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		max += int(math.Max(float64(leftDepth), float64(rightDepth)))
	}
	max++
	return max
}
