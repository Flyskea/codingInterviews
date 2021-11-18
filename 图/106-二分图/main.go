package main

import "fmt"

// BFS
func isBipartite(graph [][]int) bool {
	var (
		nodesLen = len(graph)
		colors   = make([]int, nodesLen)
		bfs      func(start, color int) bool
		queue    = make([]int, 0)
	)
	bfs = func(start, color int) bool {
		queue = append(queue, start)
		colors[start] = color
		v := 0
		for len(queue) != 0 {
			v = queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[v] {
				// 如果已染色
				if colors[neighbor] >= 1 {
					// 颜色与相邻节点相同
					if colors[neighbor] == colors[v] {
						return false
					}
				} else {
					queue = append(queue, neighbor)
					colors[neighbor] = 3 - colors[v]
				}
			}
		}
		return true
	}
	for i := 0; i < nodesLen; i++ {
		// 已染色就跳过
		if colors[i] == 0 {
			if !bfs(i, 1) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isBipartite([][]int{
		{1, 3},
		{0, 2},
		{1, 3},
		{0, 2},
	}))
}
