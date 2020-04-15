package leetcode

import (
	"math"
)

/*
{"babad", "bab"},
{"cbbd", "bb"},
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		length := int(math.Max(float64(l1), float64(l2)))
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
