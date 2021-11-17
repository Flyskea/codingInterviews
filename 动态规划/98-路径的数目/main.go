package main

import "fmt"

// dp[i][j]表示到达坐标(i, j)的路径数
// 因为只能往下或往右, 所以到达(i ,j)格的路径总数是(i-1, j)和(i, j-1)之和
// 空间优化 dp[j-1]为dp[i][j-1] dp[j]为dp[i-1][j]
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
