// https://leetcode.com/problems/move-zeroes

package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	p1 := 0
	p2 := 1
	for p2 < len(nums) {
		if nums[p1] == 0 {
			for {
				if p2 == len(nums) {
					return
				}

				if nums[p2] != 0 {
					break
				}
				p2++
			}
			nums[p1], nums[p2] = nums[p2], nums[p1]
		}

		p2++
		p1++
	}
}

func main() {
	arr1 := []int{0, 1, 2, 0, 0, 0, 3, 4}
	moveZeroes(arr1)
	fmt.Println(arr1)

	arr2 := []int{0, 0}
	moveZeroes(arr2)
	fmt.Println(arr2)
}
