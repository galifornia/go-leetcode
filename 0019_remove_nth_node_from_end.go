package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	p := head

	for p != nil {
		size++
		p = p.Next
	}

	if n < 0 || n > size {
		return head
	}

	if size < 2 {
		return &ListNode{}
	}

	if size == n {
		head = head.Next
		return head
	}

	p = head
	for i := 0; i < size-n-1; i++ {
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	} else {
		p.Next = nil
	}
	return head
}

// func main() {
// 	// six := &ListNode{Val: 6, Next: nil}
// 	// five := &ListNode{Val: 5, Next: six}
// 	// four := &ListNode{Val: 4, Next: five}
// 	// third := &ListNode{Val: 3, Next: four}
// 	second := &ListNode{Val: 2, Next: nil}
// 	head := &ListNode{Val: 1, Next: second}

// 	fmt.Println(removeNthFromEnd(head, 2))
// }
