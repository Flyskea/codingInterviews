package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(root *TreeNode, x int) bool {
	if root == nil {
		return false
	}
	if x < root.Val {
		return find(root.Left, x)
	} else if x == root.Val {
		return true
	} else {
		return find(root.Right, x)
	}
}

func findTarget(root *TreeNode, k int) bool {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node.Left != nil {
			if dfs(node.Left) {
				return true
			}
		}
		if node.Val*2 < k {
			if find(root, k-node.Val) {
				return true
			}
		} else {
			return false
		}
		if node.Right != nil {
			if dfs(node.Right) {
				return true
			}
		}
		return false
	}
	return dfs(root)
}
