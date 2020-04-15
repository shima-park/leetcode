package easy

import (
	"math"
	"strings"
)

func addBinary(a string, b string) string {
	maxLen := int(math.Max(float64(len(a)), float64(len(b))))
	a = fillZero(a, maxLen)
	b = fillZero(b, maxLen)

	var sum string
	var addBit int
	for i := maxLen - 1; i >= 0; i-- {
		av := getVal(a, i)
		bv := getVal(b, i)
		s := av + bv + addBit

		addBit = 0
		if s == 3 {
			addBit = 1
			sum = "1" + sum
		} else if s == 2 {
			addBit = 1
			sum = "0" + sum
		} else if s == 1 {
			sum = "1" + sum
		} else if s == 0 {
			sum = "0" + sum
		}
	}

	if addBit > 0 {
		sum = "1" + sum
	}

	return sum
}

func fillZero(s string, maxLen int) string {
	if len(s) < maxLen {
		s = strings.Repeat("0", maxLen-len(s)) + s
	}
	return s
}

func getVal(a string, i int) int {
	if i < len(a) {
		switch a[i] {
		case '0':
			return 0
		case '1':
			return 1
		}
	}
	return 0
}
