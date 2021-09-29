package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return walk(nums, 0, len(nums)-1)
}

func walk(nums []int, low, high int) *TreeNode {
	mid := (high-low)/2 + low
	node := &TreeNode{Val: nums[mid]}
	if low <= mid-1 {
		node.Left = walk(nums, low, mid-1)
	}

	if mid+1 <= high {
		node.Right = walk(nums, mid+1, high)
	}

	return node
}

// func main() {
// 	// [-10,-3,0,5,9]
// 	arr := []int{-10, -3, 0, 5, 9}
// 	fmt.Println(sortedArrayToBST(arr))
// }
