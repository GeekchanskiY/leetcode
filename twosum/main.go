package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i1, x := range nums {
		for i2, y := range nums {
			if i1 != i2 {
				if x+y == target {
					return []int{i1, i2}
				}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
