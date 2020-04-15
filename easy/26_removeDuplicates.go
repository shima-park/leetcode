package easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	var cursor = -1
	for i := 0; i < length; i++ {
		if cursor >= 0 && nums[i] == nums[cursor] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			length--
		} else {
			cursor++
		}
	}

	return len(nums)
}
