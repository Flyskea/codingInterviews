package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	var sum, t int
	for _, v := range nums {
		sum += v
	}
	if (sum+target)&1 == 1 || sum < target {
		return 0
	}
	t = (sum + target) / 2
	dp := make([]int, t+1)
	dp[0] = 1
	for _, v := range nums {
		for j := t; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}
	return dp[t]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}
