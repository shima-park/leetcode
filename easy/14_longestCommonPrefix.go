package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix string
	var cursor int
	for {
		var r byte
		if len(strs[0]) > cursor {
			r = strs[0][cursor]
		} else {
			break
		}

		var match = true
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= cursor || r != strs[i][cursor] {
				match = false
				break
			}
		}

		if match {
			prefix += string(r)
		} else {
			break
		}

		cursor++
	}
	return prefix
}
