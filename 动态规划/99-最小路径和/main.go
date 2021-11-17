package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[i][j]表示到达坐标(i, j)的和
// 因为只能往下或往右, 所以到达(i ,j)格的路径和是(i-1, j)或(i, j-1)加上当前格子的数字
// 空间优化 dp[j-1]为dp[i][j-1] dp[j]为dp[i-1][j]
func minPathSum(grid [][]int) int {
	var (
		colLen = len(grid[0])
		dp     = make([]int, colLen)
	)
	for i := range grid {
		for j, v := range grid[i] {
			if i == 0 && j == 0 {
				dp[j] = v
			} else if i == 0 {
				dp[j] = dp[j-1] + v
			} else if j == 0 {
				dp[j] += v
			} else {
				dp[j] = min(dp[j], dp[j-1]) + v
			}
		}
	}
	return dp[colLen-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
