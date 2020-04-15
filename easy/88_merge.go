package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && n > 0 {
		copy(nums1, nums2)
		return
	}

	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k], nums1[i] = nums1[i], nums1[k]
			i--
		}
		k--
	}

	if j >= 0 {
		copy(nums1, nums2[:j+1])
	}
}
