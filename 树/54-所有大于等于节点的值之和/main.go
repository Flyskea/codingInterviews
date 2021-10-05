package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBSTCore(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = convertBSTCore(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	root.Left = convertBSTCore(root.Left, sum)
	return root
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	return convertBSTCore(root, &sum)
}

func main() {
	fmt.Println(convertBST(nil))
}
