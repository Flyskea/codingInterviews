package main

import "fmt"

func canPartition(nums []int) bool {
	var (
		sum, target int
	)
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	target = sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
