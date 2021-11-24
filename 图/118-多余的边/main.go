package main

import (
	"fmt"
)

func Find(disjointSet []int, x int) int {
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

func findRedundantConnection(edges [][]int) []int {
	var (
		disjointSet = make([]int, len(edges)+1)
	)
	for i := range disjointSet {
		disjointSet[i] = -1
	}
	for _, v := range edges {
		node1, node2 := Find(disjointSet, v[0]), Find(disjointSet, v[1])
		if node1 == node2 {
			return []int{v[0], v[1]}
		} else {
			Union(disjointSet, node1, node2)
		}
	}
	return nil
}

func main() {
	fmt.Println(findRedundantConnection([][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}))
}
