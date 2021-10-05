package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumCore(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := max(maxPathSumCore(root.Left, maxSum), 0)
	right := max(maxPathSumCore(root.Right, maxSum), 0)
	sum := root.Val + left + right
	if *maxSum < sum {
		*maxSum = sum
	}
	return root.Val + max(left, right)
}

func maxPathSum(root *TreeNode) int {
	max := root.Val
	maxPathSumCore(root, &max)
	return max
}

func main() {
	fmt.Println(maxPathSum(nil))
}
