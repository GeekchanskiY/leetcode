package main

import (
	"fmt"
)

// speed: O(3n)
func threeSum(nums []int) [][]int {
	var (
		clean   = make([][]int, 0)
		storage = make(map[int][][]int)
	)

	for i, num := range nums {
		for j, num2 := range nums {
			if i != j {
				sum := num + num2
				if _, found := storage[sum]; found {
					found := false
					for _, pair := range storage[sum] {
						if pair[0] == i && pair[1] == j || pair[0] == j && pair[1] == i {
							found = true
							break
						}
					}
					if !found {
						storage[sum] = append(storage[sum], []int{i, j})
					}

				} else {
					storage[sum] = [][]int{{i, j}}
				}
			}
		}
	}

	for i, num := range nums {
		need := -num
		if _, found := storage[need]; found {
			fmt.Println(need, storage[need])
			for j, pair := range storage[need] {
				if pair[0] != i && pair[1] != i && len(pair) == 2 {
					storage[need][j] = append(storage[need][j], i)
					clean = append(clean, storage[need][j])
				}
			}
		}
	}

	return clean
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
}
