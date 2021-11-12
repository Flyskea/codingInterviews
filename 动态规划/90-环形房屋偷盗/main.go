package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robCore(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var (
		pre = nums[0]
		// 第[i]间房最大偷窃金额
		cur     = max(nums[0], nums[1])
		numsLen = len(nums)
	)
	for i := 2; i < numsLen; i++ {
		pre, cur = cur, max(pre+nums[i], cur)
	}
	return cur
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 忽略第一个和最后一个
	return max(robCore(nums[1:]), robCore(nums[:len(nums)-1]))
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
}
