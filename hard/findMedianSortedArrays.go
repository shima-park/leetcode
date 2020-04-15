package leetcode

import (
	"sort"
)

// 中位数是数组中的中位值
// 如果数组长度奇数则取[有序]数组的对半下标值
// 如果数组长度偶数则取[有序]数组最中间的两个值的平均值
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := append(nums1, nums2...)
	sort.Ints(merge)
	length := len(merge)
	if length%2 == 0 {
		return (float64(merge[length/2-1]) + float64(merge[length/2])) / 2
	}
	return float64(merge[length/2])
}
