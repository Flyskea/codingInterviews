package main

import "fmt"

func subsets(nums []int) [][]int {
	var (
		dfs func(values, tmp []int, start int)
		ans [][]int
	)
	dfs = func(values, tmp []int, start int) {
		ans = append(ans, append([]int{}, tmp...))
		for i := start; i < len(values); i++ {
			tmp = append(tmp, values[i])
			dfs(values, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(nums, []int{}, 0)
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
