package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		return findMin(root.Left)
	}
	return root
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root.Val == p.Val {
		return findMin(root.Right)
	}
	var ans *TreeNode
	if root.Val > p.Val {
		ans = inorderSuccessor(root.Left, p)
	} else {
		ans = inorderSuccessor(root.Right, p)
	}
	if ans == nil && root.Val > p.Val {
		ans = root
	}
	return ans
}

// 迭代
/* func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var cur, ans *TreeNode
	cur = root
	for cur != nil {
		if cur.Val > p.Val {
			ans = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return ans
} */

func main() {
	fmt.Println(inorderSuccessor(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}, &TreeNode{Val: 1}))
}
