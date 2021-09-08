package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	if head.Next == nil {
		return head
	}
	p2 := head.Next

	// remove duplicates of very first element
	for p2 != nil && p.Val == p2.Val {
		for p2 != nil && p.Val == p2.Val {
			p2 = p2.Next
		}
		if p2 == nil {
			return nil
		}

		head = p2
		p = head
		p2 = p.Next
	}

	duplicate := false
	for p2 != nil {
		if p2.Next != nil && p2.Val == p2.Next.Val {
			p2 = p2.Next
			duplicate = true
		} else if duplicate {
			p2 = p2.Next
			duplicate = false
		} else {
			p.Next = p2
			p2 = p2.Next
			p = p.Next
		}
	}
	// if p.Val == head.Val {
	p.Next = nil
	// }

	return head
}

// func main() {
// 	//  [1,1,2,2]
// 	// node7 := &ListNode{Val: 5, Next: nil}
// 	// node6 := &ListNode{Val: 4, Next: node7}
// 	// node5 := &ListNode{Val: 4, Next: node6}
// 	node4 := &ListNode{Val: 3, Next: nil}
// 	node3 := &ListNode{Val: 3, Next: node4}
// 	node2 := &ListNode{Val: 2, Next: node3}
// 	l1 := &ListNode{Val: 1, Next: node2}

// 	output := deleteDuplicates(l1)

// 	for output != nil {
// 		fmt.Printf("%d ", output.Val)
// 		output = output.Next
// 	}
// 	fmt.Println()
// }
