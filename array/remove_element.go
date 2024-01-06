package array

func removeElement(nums []int, val int) int {
	//快慢双指针思路
	slow, fast := 0, 0
	for fast < len(nums) {
		for fast < len(nums) && nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
		fast++
	}
	return slow
}
