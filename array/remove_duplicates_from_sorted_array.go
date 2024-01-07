package array

// from: https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		for fast < len(nums) && nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
		fast++
	}
	return slow + 1
}
