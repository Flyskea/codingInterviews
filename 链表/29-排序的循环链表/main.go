package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insertCore(head, node *Node) {
	var (
		cur     *Node = head
		next    *Node = head.Next
		biggest *Node = head
	)
	for !(cur.Val <= node.Val && next.Val >= node.Val) && next != head {
		cur = next
		next = next.Next
		if cur.Val > biggest.Val {
			biggest = cur
		}
	}
	if cur.Val <= node.Val && next.Val >= node.Val {
		cur.Next = node
		node.Next = next
	} else {
		node.Next = biggest.Next
		biggest.Next = node
	}
}

func insert(aNode *Node, x int) *Node {
	node := &Node{
		Val: x,
	}
	if aNode == nil {
		node.Next = node
		aNode = node
	} else if aNode.Next == nil {
		aNode.Next = node
		node.Next = aNode
	} else {
		insertCore(aNode, node)
	}
	return aNode
}

func main() {
	fmt.Println(insert(nil, 0))
}
