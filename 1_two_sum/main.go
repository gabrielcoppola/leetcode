package main

import "fmt"

func twoSum(nums []int, target int) []int {
    for idx_1, i := range nums {
		for idx_2, j := range nums {
			if idx_1 != idx_2 && (i + j) == target {
				return []int{idx_1, idx_2}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{3,3}
	result := twoSum(nums, 6)
	fmt.Println(result)
}