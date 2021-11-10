package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(l1, l2 *ListNode) *ListNode {
	var (
		head = &ListNode{}
		p    = head
	)
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		slow, fast = head, head
	)
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow.Next
	slow.Next = nil
	return mergeList(sortList(head), sortList(cur))
}

func main() {
	fmt.Println(sortList(
		&ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	))
}
