// https://leetcode.com/problems/remove-element

package main

import "fmt"

func removeElement(nums []int, val int) int {
	result := make([]int, 0, len(nums))

	for _, n := range nums {
		if n != val {
			result = append(result, n)
		}
	}

	copy(nums, result)

	return len(result)
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 5, 6, 7}
	fmt.Println(removeElement(arr, 4))
}
