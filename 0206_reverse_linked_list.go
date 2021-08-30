package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reverseList(head *ListNode) *ListNode {
	reversedList := &ListNode{}

	for head != nil {
		aux := &ListNode{Val: head.Val, Next: reversedList.Next}
		reversedList.Next = aux
		head = head.Next
	}
	return reversedList.Next
}

// func main() {
// 	node3 := &ListNode{Val: 4, Next: nil}
// 	node2 := &ListNode{Val: 2, Next: node3}
// 	l1 := &ListNode{Val: 1, Next: node2}

// 	fmt.Println(reverseList(l1))
// }
