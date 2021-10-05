package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flattenCore(root *Node) *Node {
	node := root
	var tail, next, child *Node
	for node != nil {
		next = node.Next
		if node.Child != nil {
			child = node.Child
			tail = flattenCore(node.Child)
			node.Child = nil
			node.Next = child
			child.Prev = node
			tail.Next = next
			if next != nil {
				next.Prev = tail
			}
		} else {
			tail = node
		}
		node = next
	}
	return tail
}

func flatten(root *Node) *Node {
	flattenCore(root)
	return root
}

func main() {
	fmt.Println(flatten(nil))
}
