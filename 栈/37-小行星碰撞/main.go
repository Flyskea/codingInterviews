package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, v := range asteroids {
		for len(stack) != 0 && stack[len(stack)-1] > 0 && -v > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 && v < 0 && stack[len(stack)-1] == -v {
			stack = stack[:len(stack)-1]
		} else if v > 0 || len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, v)
		}
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
}
