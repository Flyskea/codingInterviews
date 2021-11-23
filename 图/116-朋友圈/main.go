package main

import "fmt"

func Find(disjointSet []int, x int) int {
	if disjointSet[x] <= 0 {
		return x
	}
	root := x
	tmp := 0
	for disjointSet[root] >= 0 {
		root = disjointSet[root]
	}
	for x != root {
		tmp = disjointSet[x]
		disjointSet[x] = root
		x = tmp
	}
	return root
}

func Union(disjointSet []int, root1, root2 int) {
	if disjointSet[root1] > disjointSet[root2] {
		disjointSet[root2] += disjointSet[root1]
		disjointSet[root1] = root2
	} else {
		disjointSet[root1] += disjointSet[root2]
		disjointSet[root2] = root1
	}
}

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}
	var node1, node2, count int
	rows, cols := len(isConnected), len(isConnected[0])
	disjSet := make([]int, len(isConnected))
	for i := range disjSet {
		disjSet[i] = -1
	}
	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			if isConnected[i][j] == 1 {
				node1, node2 = Find(disjSet, i), Find(disjSet, j)
				if node1 != node2 {
					Union(disjSet, node1, node2)
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		if disjSet[i] < 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
}
