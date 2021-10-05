package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		cur:   root,
		stack: make([]*TreeNode, 0, 8),
	}
}

func (bst *BSTIterator) Next() int {
	for bst.cur != nil {
		bst.stack = append(bst.stack, bst.cur)
		bst.cur = bst.cur.Left
	}
	bst.cur = bst.stack[len(bst.stack)-1]
	bst.stack = bst.stack[:len(bst.stack)-1]
	val := bst.cur.Val
	bst.cur = bst.cur.Right
	return val
}

func (bst *BSTIterator) HasNext() bool {
	return bst.cur != nil || len(bst.stack) != 0
}

func main() {
	b := Constructor(&TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	})
	fmt.Println(b.Next())
	fmt.Println(b.Next())
	fmt.Println(b.HasNext())
	fmt.Println(b.Next())
	fmt.Println(b.HasNext())
	fmt.Println(b.Next())
	fmt.Println(b.HasNext())
	fmt.Println(b.Next())
	fmt.Println(b.HasNext())
}
