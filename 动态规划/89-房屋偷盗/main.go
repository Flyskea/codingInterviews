package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
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

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
