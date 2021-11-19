package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		dfs func(int)
		cur = make([]int, 1)
		ans [][]int
	)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int{}, cur...))
			return
		}
		for _, y := range graph[x] {
			cur = append(cur, y)
			dfs(y)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
