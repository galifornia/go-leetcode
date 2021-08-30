package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	head := mergedList

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			mergedList.Next = l1
			l1 = l1.Next
		} else {
			mergedList.Next = l2
			l2 = l2.Next
		}
		mergedList = mergedList.Next
	}

	if l1 != nil {
		mergedList.Next = l1
	} else {
		mergedList.Next = l2
	}

	return head.Next
}

// func main() {
// 	node3 := &ListNode{Val: 4, Next: nil}
// 	node2 := &ListNode{Val: 2, Next: node3}
// 	l1 := &ListNode{Val: 1, Next: node2}

// 	node6 := &ListNode{Val: 4, Next: nil}
// 	node5 := &ListNode{Val: 3, Next: node6}
// 	l2 := &ListNode{Val: 1, Next: node5}

// 	fmt.Println(mergeTwoLists(l1, l2))
// }
