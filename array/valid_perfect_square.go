package array

// from: https://leetcode.cn/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	// 转换成 k^2 = x 判断是否有满足该条件的k值
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
