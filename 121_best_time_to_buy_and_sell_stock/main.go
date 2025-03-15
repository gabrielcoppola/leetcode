package main

import "fmt"
import "math"

func maxProfit(prices []int) int {
	small := math.MaxInt64
	profit := 0
	for i := 0; i < len(prices); i++ {
		if small > prices[i] {
			small = prices[i]
		}
		if profit < (prices[i] - small) {
			profit = prices[i] - small
		}
	}
	return profit
}

func main() {
	prices := []int{7,1,5,3,6,4}
	result := maxProfit(prices)
	fmt.Println(result)
}