package main

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	solution := []int{}

	traverse(root, &solution)

	return solution
}

func traverse(node *TreeNode, solution *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		traverse(node.Left, solution)
	}

	*solution = append(*solution, node.Val)

	if node.Right != nil {
		traverse(node.Right, solution)
	}
}
