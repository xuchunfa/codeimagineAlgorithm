package array

// from: https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

// 时间复杂度O(log n)
func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := leftBorder(nums, target), rightBorder(nums, target)
	if left == -2 || right == -2 {
		return []int{-1, -1}
	}
	if right-left > 1 {
		return []int{left + 1, right - 1}
	}
	return []int{-1, -1} // [1,3,5] target=2的情况
}

func leftBorder(nums []int, target int) int {
	leftBorder := -2              // [1,2,3] target=4 的情况
	left, right := 0, len(nums)-1 // [left, right]区间划分思路
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			right = mid - 1
			leftBorder = right
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

func rightBorder(nums []int, target int) int {
	rightBorder := -2             // [1,2,3] target=0 的情况
	left, right := 0, len(nums)-1 // [left, right]区间划分思路
	for left <= right {
		mid := left + (right-left)>>1
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
			rightBorder = left
		}
	}
	return rightBorder
}
