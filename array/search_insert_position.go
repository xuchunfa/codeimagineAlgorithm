package array

// from: https://leetcode.cn/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1 // [left, right]区间划分思路
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
