package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var (
		dfs func(t, start int)
		ans [][]int
		cur []int
	)
	dfs = func(t, start int) {
		if t <= 0 {
			if t == 0 {
				ans = append(ans, append([]int{}, cur...))
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(t-candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return ans
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
