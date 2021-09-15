package main

import "fmt"

type Path struct {
	last  int
	route []int
}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}
	res := [][]int{}
	paths = append(paths, []int{0})
	for len(paths) > 0 {
		p := paths[0]
		paths = paths[1:]
		// last element in path
		le := p[len(p)-1]
		// check if last element is end of our path
		if le == len(graph)-1 {
			res = append(res, p)
			continue
		}
		// add all posible path branches
		for _, n := range graph[le] {
			np := make([]int, len(p)+1)
			copy(np, p)
			np[len(p)] = n
			paths = append(paths, np)
		}
	}
	return res
}

func main() {
	// arr := [][]int{{1, 2}, {3}, {3}, {}}
	arr := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(arr))
}
