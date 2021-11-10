package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var (
		ans [][]int
		dfs func(target, start int, cur []int)
	)
	dfs = func(target, start int, cur []int) {
		if target <= 0 {
			if target == 0 {
				ans = append(ans, append([]int{}, cur...))
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			dfs(target-candidates[i], i, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0, []int{})
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
