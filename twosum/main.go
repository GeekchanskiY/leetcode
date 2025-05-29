package main

import "fmt"

func twoSumHashMap(nums []int, target int) []int {
	res := make([]int, 2)
	diffMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if j, found := diffMap[diff]; found {
			res[0] = j
			res[1] = i

			return res
		} else {
			diffMap[nums[i]] = i
		}
	}

	return res
}

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
	fmt.Println(twoSumHashMap(nums, target))
}
