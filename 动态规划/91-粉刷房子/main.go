package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(costs [][]int) int {
	var (
		costsLen = len(costs)
		dp       = make([][3]int, costsLen)
	)

	for i := 0; i < 3; i++ {
		dp[0][i] = costs[0][i]
	}
	for i := 1; i < costsLen; i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + costs[i][2]
	}
	return min(dp[costsLen-1][0], min(dp[costsLen-1][1], dp[costsLen-1][2]))
}

func main() {
	fmt.Println(minCost([][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}))
}
