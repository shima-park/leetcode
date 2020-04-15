package easy

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	var index = -1
	if len(needle) > len(haystack) {
		return index
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			matched := true
			for j := 0; j < len(needle); j++ {
				if j+i > len(haystack)-1 || haystack[j+i] != needle[j] {
					matched = false
					break
				}
			}
			if matched {
				index = i
				break
			}
		}
	}

	return index
}
