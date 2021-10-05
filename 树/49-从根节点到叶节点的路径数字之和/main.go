package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbersCore(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return sumNumbersCore(root.Left, sum) + sumNumbersCore(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersCore(root, 0)
}

func main() {
	fmt.Println(sumNumbers(nil))
}
