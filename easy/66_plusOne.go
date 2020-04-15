package easy

func plusOne(digits []int) []int {
	var add = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+add%10 == 0 {
			digits[i] = 0
			add = 1
		} else {
			if add > 0 {
				digits[i] = digits[i] + add
				add = 0
			}
		}
	}
	return digits
}
