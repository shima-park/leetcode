package leetcode

func maxArea(height []int) int {
	var maxarea int
	var l int
	var r = len(height) - 1
	for l < r {
		maxarea = max(
			maxarea,
			(r-l)*min(height[l], height[r]),
		)

		if l < r {
			l++
		} else {
			r--
		}
	}
	return maxarea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
