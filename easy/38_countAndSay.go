package easy

import "fmt"

/*
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var s = "1"
	var ret string
	for i := 1; i < n; i++ {
		var c int
		var b byte
		ret = ""
		for k := 0; k < len(s); k++ {
			if b != s[k] {
				if c > 0 {
					ret += fmt.Sprint(c) + string(b)
				}
				b = s[k]
				c = 1
			} else {
				c++
			}
		}
		ret += fmt.Sprint(c) + string(b)
		s = ret
	}

	return ret
}
