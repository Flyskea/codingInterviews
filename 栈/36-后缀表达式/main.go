package main

import (
	"fmt"
	"strconv"
)

func calculate(num1, num2 int, op string) int {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := []int{}
	var (
		num1, num2, tmp int
	)
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			num1 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, calculate(num2, num1, token))
		default:
			tmp, _ = strconv.Atoi(token)
			stack = append(stack, tmp)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
