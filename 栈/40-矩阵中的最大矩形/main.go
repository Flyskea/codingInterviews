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

func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for _, row := range matrix {
		for i, v := range row {
			if v == '0' {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		maxArea = max(maxArea, largestRectangleArea(heights))
	}
	return maxArea
}

func main() {
	fmt.Println(maximalRectangle([]string{"10100", "10111", "11111", "10010"}))
}
