package main

import (
	"fmt"
	"math"
)

func AbsInt(x int64) int64 {
	if x > 0 {
		return x
	}
	return x * -1
}

func findClosestNumber(nums []int) int {
	close := int64(math.MaxInt64)
	for i := 0; i < len(nums); i++ {
		num := int64(nums[i])
		if AbsInt(num) < AbsInt(close) || (AbsInt(num) == AbsInt(close) && num > close) {
			close = num
		}
	}
	return int(close)
}

func main() {
	nums := []int{2,1,-1}
	res := findClosestNumber(nums)
	fmt.Println(res)
}