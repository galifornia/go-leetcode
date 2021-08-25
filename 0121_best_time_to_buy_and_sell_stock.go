package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else {
			currentProfit := v - minPrice
			if currentProfit > profit {
				profit = currentProfit
			}
		}
	}
	return profit
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{1, 2}
	fmt.Println(maxProfit(prices))
}
