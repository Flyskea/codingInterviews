package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	var (
		ans    = math.MaxInt32
		rowLen = len(triangle)
		dp     = make([][]int, rowLen)
	)
	for i := range dp {
		dp[i] = make([]int, rowLen)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < rowLen; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = triangle[i][j]
			if j == 0 {
				dp[i][j] += dp[i-1][j]
			} else if i == j {
				dp[i][j] += dp[i-1][j-1]
			} else {
				dp[i][j] += min(dp[i-1][j], dp[i-1][j-1])
			}
			if i == rowLen-1 {
				ans = min(ans, dp[i][j])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
