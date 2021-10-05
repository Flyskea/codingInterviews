package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	index := 0
	for i, v := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < v {
			index = stack[len(stack)-1]
			ans[index] = i - index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
