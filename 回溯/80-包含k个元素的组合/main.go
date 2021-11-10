package main

import "fmt"

func combine(n int, k int) [][]int {
	var (
		dfs func(int)
		ans [][]int
		cur []int
	)
	dfs = func(index int) {
		if len(cur) == k {
			ans = append(ans, append([]int{}, cur...))
		}
		if len(cur)+(n-index+1) < k {
			return
		}
		for i := index; i <= n; i++ {
			cur = append(cur, i)
			dfs(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(1)
	return ans
}

func main() {
	fmt.Println(combine(3, 1))
}
