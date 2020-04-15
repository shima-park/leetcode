package leetcode

import (
	"math"
)

func myAtoi(str string) int {
	var rev int
	var isNegative bool
	var isStatedConvert bool
	for i := 0; i < len(str); i++ {
		// 空格 conitnue
		if str[i] == 32 && !isStatedConvert {
			continue
		}

		if isNum(str[i]) {
			if !isStatedConvert {
				isStatedConvert = true
			}
			num := int(str[i] % 48)

			if isNegative {
				num *= -1
			}

			if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && num > 7) {
				return math.MaxInt32
			}

			if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && num < -8) {
				return math.MinInt32
			}

			rev = rev*10 + num
		} else {
			if str[i] != 43 && str[i] != 45 {
				break
			}

			if isStatedConvert {
				break
			} else {
				if str[i] == 45 {
					isNegative = true
				}
			}

			// 因为它的下一个字符不为数字 break
			if !checkNum(str, i+1) {
				break
			}
		}
	}

	return rev
}

func checkNum(s string, i int) bool {
	if i >= 0 && i < len(s) {
		return isNum(s[i])
	}
	return false
}

func isNum(b byte) bool {
	return b >= 48 && b <= 57
}
