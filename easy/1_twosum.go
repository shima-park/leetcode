package easy

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for i, num := range nums {
		i2, ok := m[target-num]
		if ok && i != i2 {
			return []int{i, i2}
		}
		m[num] = i
	}

	return nil
}
