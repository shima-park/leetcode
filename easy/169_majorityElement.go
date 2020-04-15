package easy

func majorityElement(nums []int) int {
	var maxNum int
	var m = map[int]int{}
	for _, n := range nums {
		m[n]++
		if m[n] > m[maxNum] {
			maxNum = n
		}
	}

	return maxNum
}
