package main

import "fmt"

func permute(nums []int) [][]int {
	var (
		ans [][]int
		dfs func(int)
	)
	dfs = func(start int) {
		if start == len(nums)-1 {
			ans = append(ans, append([]int{}, nums...))
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			dfs(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
