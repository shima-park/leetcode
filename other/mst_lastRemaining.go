package leetcode

func lastRemaining(n int, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	var c int
	for len(arr) == 1 {
		delEle(arr, (c+m)%m)
		c += m
	}

	return arr[0]
}

func delEle(arr []int, i int) {
	arr = append(arr[:i], arr[i+1:]...)
}
