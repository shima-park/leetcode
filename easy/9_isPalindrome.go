package easy

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var rev int
	for x > rev { // 判断是否反转到数字作为字符串来看时一半的位置
		rev = rev*10 + x%10
		x /= 10
	}
	fmt.Println(x, rev, rev/10)
	return x == rev || x == rev/10
}
