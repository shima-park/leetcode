package easy

func lengthOfLastWord(s string) int {
	var c int
	var started bool
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 64 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			if !started {
				started = true
			}
			c++
		} else {
			if started {
				break
			}
		}
	}

	return c
}
