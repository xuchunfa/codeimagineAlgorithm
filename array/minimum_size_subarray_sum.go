package array

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// from: https://leetcode.cn/problems/minimum-size-subarray-sum/submissions/494391565/
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口算法
	var res = math.MaxInt32
	var sum, start int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target { //start坐标赶上i坐标,此时sum=0
			res = Min(res, i-start+1)
			sum -= nums[start]
			start++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
