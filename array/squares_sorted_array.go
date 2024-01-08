package array

func SortedSquares(nums []int) []int {
	// 由于存在负数,最大值在有序数组的两头 -> O(n)时间复杂度解题思路
	if nums == nil {
		return []int{}
	}
	//注意和res := make([]int, 0, len(nums))的区别
	res := make([]int, len(nums)) //从后往前填充数组的最大值
	k := len(nums) - 1
	for i, j := 0, len(nums)-1; i < len(nums) && j >= 0 && i <= j; {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return res
}
