package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var (
		ans     [][]int
		cur     []int
		dfs     func([]int)
		used    = make([]bool, len(nums))
		numsLen = len(nums)
	)
	dfs = func(cur []int) {
		if len(cur) == numsLen {
			ans = append(ans, append([]int{}, cur...))
		}
		for i := 0; i < numsLen; i++ {
			if !used[i] {
				if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
					continue
				}
				used[i] = true
				cur = append(cur, nums[i])
				dfs(cur)
				cur = cur[:len(cur)-1]
				used[i] = false
			}
		}
	}
	dfs(cur)
	return ans
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
