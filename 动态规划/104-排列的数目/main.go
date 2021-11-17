package main

import "fmt"

// 用 f(i) 表示和为 i 的排列数目。为了得到和为 i 的排列，有如下选择：
// 在和为 i - nums[0] 的排列中添加下标为 0 的数字，此时 f(i) 等于 f(i - nums[0])
// 依次类推，在和为 i - nums[1] 的排列中添加下标为 n - 1 的数字，此时 f(i) 等于 f(i - nums[n - 1])
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
