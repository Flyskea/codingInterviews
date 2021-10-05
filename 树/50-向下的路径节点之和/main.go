package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSumCore(root *TreeNode, targetSum, sum int, sumMap map[int]int) int {
	if root == nil {
		return 0
	}
	sum += root.Val
	count := 0
	count += sumMap[sum-targetSum]
	sumMap[sum]++
	count += pathSumCore(root.Left, targetSum, sum, sumMap)
	count += pathSumCore(root.Right, targetSum, sum, sumMap)
	sumMap[sum]--
	return count
}

func pathSum(root *TreeNode, targetSum int) int {
	return pathSumCore(root, targetSum, 0, map[int]int{0: 1})
}

func main() {
	fmt.Println(pathSum(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}, 3))
}
