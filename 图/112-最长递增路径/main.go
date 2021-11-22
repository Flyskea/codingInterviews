package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	var (
		rows        = len(matrix)
		cols        = len(matrix[0])
		dfs         func(row, col int) int
		dx          = []int{0, -1, 1, 0}
		dy          = []int{1, 0, 0, -1}
		max, length int
		lengths     = make([][]int, rows)
	)
	for i := range lengths {
		lengths[i] = make([]int, cols)
	}
	dfs = func(row, col int) int {
		if lengths[row][col] != 0 {
			return lengths[row][col]
		}
		cur := 1
		for i := 0; i < 4; i++ {
			x, y := row+dx[i], col+dy[i]
			if x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] > matrix[row][col] {
				path := dfs(x, y) + 1
				if path > cur {
					cur = path
				}
			}
		}
		lengths[row][col] = cur
		return cur
	}
	for i := range matrix {
		for j := range matrix[i] {
			length = dfs(i, j)
			if length > max {
				max = length
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestIncreasingPath([][]int{
		{1, 2},
		{2, 3},
	}))
}
