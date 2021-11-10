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

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeList(left, right)
}

func main() {
	a := []*ListNode{}
	a = append(a, &ListNode{
		Val: 3,
	})
	a = append(a, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	})
	fmt.Println(mergeKLists(a))
}
