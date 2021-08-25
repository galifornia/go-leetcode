package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	size := 0
	p := head
	for p != nil {
		p = p.Next
		size++
	}

	p = head
	for i := 0; i < size/2; i++ {
		p = p.Next
	}
	return p
}

// func main() {
// 	six := &ListNode{Val: 6, Next: nil}
// 	five := &ListNode{Val: 5, Next: six}
// 	four := &ListNode{Val: 4, Next: five}
// 	third := &ListNode{Val: 3, Next: four}
// 	second := &ListNode{Val: 2, Next: third}
// 	head := &ListNode{Val: 1, Next: second}

// 	fmt.Println(middleNode(head))
// }
