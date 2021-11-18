package main

import (
	"fmt"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	var (
		rows       = len(mat)
		cols       = len(mat[0])
		dists      = make([][]int, rows)
		queue      = make([][2]int, 0)
		dx         = []int{0, 1, -1, 0}
		dy         = []int{1, 0, 0, -1}
		cur        [2]int
		dist, x, y int
	)
	for i := 0; i < rows; i++ {
		dists[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				dists[i][j] = 0
			} else {
				dists[i][j] = math.MaxInt32
			}
		}
	}
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		dist = dists[cur[0]][cur[1]]
		for i := 0; i < 4; i++ { //遍历当前坐标的上下左右
			x, y = cur[0]+dx[i], cur[1]+dy[i]
			if x >= 0 && x < rows && y >= 0 && y < cols && dists[x][y] > dist+1 {
				dists[x][y] = dist + 1 //更新坐标，和相邻的坐标+1
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	return dists
}

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
