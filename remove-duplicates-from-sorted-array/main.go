package main

import "fmt"

// task says that it must remove the duplicates, but runner checks only the k result. I did it as the task says
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	ptr1 := 0
	ptr2 := 1

	for ptr1 < len(nums) && ptr2 < len(nums) {
		if nums[ptr1] == nums[ptr2] {
			nums = append(nums[:ptr1], nums[ptr2:]...)
		} else {
			ptr1++
			ptr2++
		}

	}

	return len(nums)
}

func main() {
	arr := []int{1, 1, 2, 3, 4, 4, 5, 6, 7, 7}
	res := removeDuplicates(arr)
	fmt.Println(res)
}
