package main

import "fmt"

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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var (
		levelLen, maxLevel int
		ans                []int
		v                  *TreeNode
	)
	for len(queue) != 0 {
		levelLen = len(queue)
		maxLevel = queue[0].Val
		for i := 0; i < levelLen; i++ {
			v = queue[0]
			queue = queue[1:]
			maxLevel = max(maxLevel, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		ans = append(ans, maxLevel)
	}
	return ans
}

func main() {
	fmt.Println(largestValues(&TreeNode{
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
