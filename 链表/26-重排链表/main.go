package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeList(l1 *ListNode, l2 *ListNode) {
	var l1tmp, l2tmp *ListNode
	for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next
		l1.Next = l2
		l1 = l1tmp
		l2.Next = l1
		l2 = l2tmp
	}
}

func reverseList(head *ListNode) *ListNode {
	var cur, next, pre *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	l1 := head
	tmp = reverseList(tmp)
	mergeList(l1, tmp)
}

func main() {
	left := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	reorderList(left)
	fmt.Println(left)
}
