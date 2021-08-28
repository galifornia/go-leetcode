package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root != nil {
		assignRightPointer(root.Left, root.Right)
	}

	return root
}

func assignRightPointer(node1, node2 *Node) {
	if node1 != nil {
		node1.Next = node2
		assignRightPointer(node1.Left, node1.Right)
		assignRightPointer(node1.Right, node2.Left)
		if node2 != nil {
			assignRightPointer(node2.Left, node2.Right)
		}
	}
}

func main() {
	node7 := &Node{Val: 7}
	node6 := &Node{Val: 6}
	node5 := &Node{Val: 5}
	node4 := &Node{Val: 4}
	node3 := &Node{Val: 3, Left: node6, Right: node7}
	node2 := &Node{Val: 2, Left: node4, Right: node5}
	node1 := &Node{Val: 1, Left: node2, Right: node3}

	fmt.Println(connect((node1)))
}
