package easy

var mapping = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		left, isRight := mapping[s[i]]
		if !isRight {
			stack = append(stack, s[i])
		} else {
			stackLen := len(stack)
			if stackLen == 0 {
				return false
			}

			if stack[stackLen-1] != left {
				return false
			}

			stack = stack[:stackLen-1]
		}
	}

	return len(stack) == 0
}
