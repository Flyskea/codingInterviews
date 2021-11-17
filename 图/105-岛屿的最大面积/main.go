package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxAreaOfIsland(grid [][]int) int {
	var (
		rows    = len(grid)
		cols    = len(grid[0])
		maxArea int
		dfs     func(row, col int) int
		dx      []int = []int{0, 1, -1, 0}
		dy      []int = []int{1, 0, 0, -1}
	)
	dfs = func(row, col int) int {
		if grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0
		var x, y, area int
		area = 1
		for i := 0; i < 4; i++ {
			x = row + dx[i]
			y = col + dy[i]
			if x >= 0 && x < rows && y >= 0 && y < cols {
				area += dfs(x, y)
			}
		}
		return area
	}
	if rows == 0 || cols == 0 {
		return 0
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{}))
}
