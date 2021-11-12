package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	pre, cur := cost[0], cost[1]
	for i := 2; i < n; i++ {
		pre, cur = cur, min(cur, pre)+cost[i]
	}
	return min(pre, cur)
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
