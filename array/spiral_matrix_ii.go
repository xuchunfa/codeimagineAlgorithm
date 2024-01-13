package array

// from: https://leetcode.cn/problems/spiral-matrix-ii/description/

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	var k = 1
	top, bottom, left, right := 0, n-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			res[top][i] = k
			k++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[i][right] = k
			k++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res[bottom][i] = k
			k++
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res[i][left] = k
			k++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
