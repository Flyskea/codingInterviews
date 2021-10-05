package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var (
		levelLen int
		ans      []int
		v        *TreeNode
	)
	for len(queue) != 0 {
		levelLen = len(queue)
		for i := 0; i < levelLen; i++ {
			v = queue[0]
			queue = queue[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		ans = append(ans, v.Val)
	}
	return ans
}

func main() {
	fmt.Println(rightSideView(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}))
}
