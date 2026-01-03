// https://leetcode.com/problems/n-repeated-element-in-size-2n-array

package main

import "fmt"

func repeatedNTimes(nums []int) int {
	// Boyer-Moore algorithm would be better here

	found := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := found[nums[i]]; ok {
			return nums[i]
		}

		found[nums[i]] = struct{}{}
	}

	return 0
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}
