package leetcode

import "fmt"

func compressString(s string) string {
	var cs string
	var count int
	var char string
	var lastChar string
	for i := 0; i < len(s); i++ {
		if char != string(s[i]) {
			if count > 0 {
				cs += fmt.Sprintf("%s%d", char, count)
			}
			lastChar = char
			char = string(s[i])
			count = 0
		}
		count++
		if lastChar != char && i == len(s)-1 {
			cs += fmt.Sprintf("%s%d", char, count)
		}
	}

	if len(cs) < len(s) {
		return cs
	}
	return s
}
