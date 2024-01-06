package array

// from: https://leetcode.cn/problems/sqrtx/description/
func mySqrt(x int) int {
	var res int
	// 转换成 k^2 <= x 满足该条件的最大k值域
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
