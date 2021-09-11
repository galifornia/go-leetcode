package main

import "math"

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
    queue := make([]*Node, 0) //BFS queue
    queue = append(queue, root)

    for len(queue) > 0 {
        l := len(queue)
        for i := 0; i < l; i++ {  //iterate the length of queue only
            node := queue[0]
            queue = queue[1:]
            if i < l - 1 {
                node.Next = queue[0] //not the last node in this layer
            } else if i == l - 1 {
                node.Next = nil //the last node in this layer
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return root
} func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    queue := make([]*Node, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        l := len(queue)
        for i := 0; i < l; i++ {
            node := queue[0]
            queue = queue[1:]
            if i < l - 1 {
                node.Next = queue[0]
            } else if i == l - 1 {
                node.Next = nil
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    return root
}
