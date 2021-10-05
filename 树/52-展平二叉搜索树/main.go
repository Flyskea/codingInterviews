package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var pre *TreeNode
	var first *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Left = nil
		if pre != nil {
			pre.Right = root
			pre = pre.Right
		} else {
			first = root
			pre = root
		}
		dfs(root.Right)
	}
	dfs(root)
	return first
}

func order(root *TreeNode) {
	if root == nil {
		return
	}
	order(root.Left)
	fmt.Printf("%d ", root.Val)
	order(root.Right)
}

func main() {
	order(increasingBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 7,
		},
	}))
}
