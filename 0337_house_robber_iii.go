package main

type TreeNodeee struct {
	Val   int
	Left  *TreeNodeee
	Right *TreeNodeee
}

func rob(root *TreeNodeee) int {
	// reward of each level
	levelMap := make(map[int]int)

	findTreasure(root, &levelMap, 0)

	a, b := 0, 0

	for i := 0; i < len(levelMap); i++ {
		if i%2 == 0 {
			a = max(a+levelMap[i], b)
		} else {
			b = max(a, b+levelMap[i])
		}
	}

	return max(a, b)
}

func findTreasure(node *TreeNodeee, levelMap *map[int]int, level int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		findTreasure(node.Left, levelMap, level+1)
	}

	(*levelMap)[level] += node.Val

	if node.Right != nil {
		findTreasure(node.Right, levelMap, level+1)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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
