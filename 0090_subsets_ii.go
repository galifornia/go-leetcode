package main

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{}) // first empty element

	countNums := make(map[int]int)

	// we loop through each number in array & add it to each subarray contained in result
	for _, num := range nums {
		for _, subarray := range result {
			// count number of appeareances of num in subarray & only add it
			// if the tally is the same as the one in map countNums
			// so we do not add duplicates
			appeareances := count(subarray, num)

			if appeareances == countNums[num] {
				// if we do not create a new slice weird things happen
				// cannot do append(subarray, num)
				newSubset := append([]int{}, subarray...)
				newSubset = append(newSubset, num)
				result = append(result, newSubset)
			}
		}
		countNums[num]++
	}

	return result
}

func count(arr []int, target int) int {
	sum := 0
	for _, num := range arr {
		if num == target {
			sum++
		}
	}
	return sum
}

func main() {

}
