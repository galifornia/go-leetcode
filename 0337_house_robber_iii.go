package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return max(robber(root))
}

func robber(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	gold := node.Val
	l1, l2 := robber(node.Left)
	r1, r2 := robber(node.Right)
	c1 := gold + l2 + r2
	c2 := max(l1, l2) + max(r1, r2)
	return c1, c2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	// [3,2,3,null,3,null,1]
// 	//[2,1,3,null,4]
// 	node0 := &TreeNodeee{Val: 1}
// 	node2 := &TreeNodeee{Val: 3}
// 	node4 := &TreeNodeee{Val: 3, Right: node0}
// 	node5 := &TreeNodeee{Val: 2, Right: node2}
// 	root := &TreeNodeee{Val: 2, Left: node5, Right: node4}
// 	fmt.Println(rob(root))
// }
