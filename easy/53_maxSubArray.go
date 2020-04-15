package easy

import "math"

func maxSubArray(nums []int) int {
	var sum int
	var ans int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		ans = int(math.Max(float64(ans), float64(sum)))
	}
	return ans
}
