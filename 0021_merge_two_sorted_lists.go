package main

import "fmt"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	p := l1
	p2 := l2

	var p3 *ListNode
	var p4 *ListNode

	for p != nil || p2 != nil {
		if p != nil && p2 != nil {
			if p.Val < p2.Val {
				if p3 == nil {
					p3 = &ListNode{Val: p.Val, Next: nil}
					p4 = p3
				} else {
					next := &ListNode{Val: p.Val, Next: nil}
					p4.Next = next
					p4 = p4.Next
				}
				p = p.Next
			} else {
				if p3 == nil {
					p3 = &ListNode{Val: p2.Val, Next: nil}
					p4 = p3
				} else {
					next := &ListNode{Val: p2.Val, Next: nil}
					p4.Next = next
					p4 = p4.Next
				}
				p2 = p2.Next
			}
		} else if p != nil {
			if p3 == nil {
				p3 = &ListNode{Val: p.Val, Next: nil}
				p4 = p3
			} else {
				next := &ListNode{Val: p.Val, Next: nil}
				p4.Next = next
				p4 = p4.Next
			}
			p = p.Next
		} else if p2 != nil {
			if p3 == nil {
				p3 = &ListNode{Val: p2.Val, Next: nil}
				p4 = p3
			} else {
				next := &ListNode{Val: p2.Val, Next: nil}
				p4.Next = next
				p4 = p4.Next
			}
			p2 = p2.Next
		}
	}

	return p3
}

func main() {
	node3 := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	l1 := &ListNode{Val: 1, Next: node2}

	node6 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 3, Next: node6}
	l2 := &ListNode{Val: 1, Next: node5}

	fmt.Println(mergeTwoLists(l1, l2))
}
