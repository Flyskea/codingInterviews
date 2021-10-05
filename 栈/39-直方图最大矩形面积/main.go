package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	var curHeight, curWidth, ans int
	newHeights := make([]int, len(heights)+2)
	stack := make([]int, 0)
	for i := 1; i < len(heights)+1; i++ {
		newHeights[i] = heights[i-1]
	}
	for i, v := range newHeights {
		for len(stack) != 0 && v < newHeights[stack[len(stack)-1]] {
			curHeight = newHeights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			curWidth = i - stack[len(stack)-1] - 1
			ans = max(ans, curHeight*curWidth)
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 4}))
}
