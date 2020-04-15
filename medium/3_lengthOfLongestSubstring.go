package leetcode

func lengthOfLongestSubstring(s string) int {
	window := map[rune]int{} // key: rune value: index
	var index, max int
	for index < len(s) {
		r := rune(s[index])

		if runeIndex, found := window[r]; found {
			index = runeIndex + 1
			r = rune(s[index])

			window = map[rune]int{
				r: index,
			}
		} else {
			window[r] = index
		}

		length := len(window)
		if length > max {
			max = length
		}

		index++
	}

	return max
}
