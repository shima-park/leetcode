package leetcode

import "math"

// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
func isMatch(s string, p string) bool {
	if p == "" || s != "" && p == "" {
		return false
	}

	var pChar byte
	var pCount int
	var cursor int
	var i int
	for i < len(p) {
		// init pattern char
		pChar = p[i]

		// init pattern count
		pCount = 1
		if isMatchAny(p, i+1) { // get next char to check if any char match char
			pCount = math.MaxInt32
		}

		// {"bbbba", ".*a*a", true},
		for cursor < len(s) && pCount > 0 &&
			(s[cursor] == pChar || pChar == 46) {

			// 判断当前是否是任意匹配模式，且字符串已经即将匹配完，而p还有模式未匹配
			// 例如 s:aaa p:a*a
			if isMatchAny(p, i+1) &&
				cursor == len(s)-1 && i+2 != len(p) {
				break
			} else {
				pCount--
				cursor++
			}
		}

		// clean pattern count, when current char is *
		if isMatchAny(p, i+1) {
			pCount = 0
			i++
		} else {
			if pCount > 0 {
				return false
			}
		}

		i++
	}

	return cursor == len(s) && i == len(p) && pCount == 0
}

func isMatchAny(p string, i int) bool {
	return i >= 0 && len(p) > i && p[i] == 42
}
