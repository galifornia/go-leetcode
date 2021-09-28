package main

type TreeNodee struct {
	Val   int
	Left  *TreeNodee
	Right *TreeNodee
}

func hasPathSum(root *TreeNodee, targetSum int) bool {
	return findPath(root, targetSum, 0)
}

func findPath(node *TreeNodee, targetSum int, acc int) bool {
	if node == nil {
		return false
	}
	acc += node.Val

	// make sure its a leaf
	if acc == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	return findPath(node.Left, targetSum, acc) || findPath(node.Right, targetSum, acc)
}

// func main() {
// 	node1 := &TreeNodee{}
// 	node2 := &TreeNodee{Val: -3}
// 	root := &TreeNodee{Val: -2, Left: node1, Right: node2}
// 	target := -5
// 	fmt.Println(hasPathSum(root, target))
// }
