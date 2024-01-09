package array

// from: https://leetcode.cn/problems/fruit-into-baskets/description/
func totalFruit(fruits []int) int {
	// 巧用hash表去控制滑动窗口的左右边界
	hash := make(map[int]int)
	var start, res int
	for end, v := range fruits {
		hash[v]++
		for len(hash) > 2 { //滑动窗口内含有2类元素
			val := fruits[start]
			hash[val]--
			if hash[val] == 0 {
				delete(hash, val)
			}
			start++
		}
		res = Max(res, end-start+1)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
