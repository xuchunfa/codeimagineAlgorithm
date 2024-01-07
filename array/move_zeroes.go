package array

// from: https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	//快慢双指针思路
	//slow, fast := 0, 0
	//for fast < len(nums) {
	//	for fast < len(nums) && nums[fast] != 0 {
	//		nums[slow] = nums[fast]
	//		slow++
	//		fast++
	//	}
	//	fast++
	//}
	//for ; slow < len(nums); slow++ {
	//	nums[slow] = 0
	//}

	// 巧妙写法
	slow := 0
	for i := range nums {
		if nums[i] != 0 {
			//0被交换到后面去了
			nums[slow], nums[i] = nums[i], nums[slow]
			slow++
		}
	}
}
