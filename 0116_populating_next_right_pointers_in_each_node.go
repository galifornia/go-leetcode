package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	nodes := make([]*Node, 0)

	for len(queue) > 0 {
		node := queue[0]
		nodes = append(nodes, node)
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		nLength := len(nodes)
		// if we are on second level or lower & last appended node is on the same level to the previous
		// then point the previous one (Next) to the new one
		if nLength > 1 && math.Trunc(math.Log2(float64(nLength-1))) == math.Trunc(math.Log2(float64(nLength))) {
			nodes[len(nodes)-2].Next = nodes[nLength-1]
		}
	}

	return root
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
