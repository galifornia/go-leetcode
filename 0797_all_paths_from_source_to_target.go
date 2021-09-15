package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{0})

	for i := 0; i < len(graph[0]); i++ {
		path := make([]int, 0)
		path = append(path, 0)
		result = append(result, findPath(graph, graph[0][i], path))
	}

	return result
}

func findPath(graph [][]int, row int, path []int) []int {
	path = append(path, row)
	if row == len(graph)-1 {
		return path
	}
	for _, i := range graph[row] {
		return findPath(graph, i, path)
	}
	return path
}

func main() {
	// arr := [][]int{{1, 2}, {3}, {3}, {}}
	arr := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(arr))
}
