package easy

func searchInsert(nums []int, target int) int {
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			index++
		} else if nums[i] > target {
			break
		}
	}

	return index
}
