package leetcode

import (
	"strings"
)

/*
"LEETCODEISHIRING",
3,
`LCIRETOESIIGEDHN`,
*/
func convert(s string, n int) string {
	if n == 1 {
		return s
	}
	// 0, n-1
	ss := make([]string, n)
	curRow := 0
	goingDown := false
	for i := 0; i < len(s); i++ {
		ss[curRow] += string(s[i])
		if curRow == 0 || curRow == n-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}

	}

	var b strings.Builder
	for _, s := range ss {
		b.WriteString(s)
	}

	return b.String()
}
