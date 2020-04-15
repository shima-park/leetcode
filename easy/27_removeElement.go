package easy

func removeElement(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			length--
		}
	}
	return len(nums)
}
