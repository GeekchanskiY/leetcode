package main

import (
	"fmt"
	"sort"
)

// I am re-sorting this array, which is unnecessary.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged_arr := nums1
	for _, el := range nums2 {
		merged_arr = append(merged_arr, el)
	}

	sort.Slice(merged_arr, func(i, j int) bool {
		return merged_arr[i] > merged_arr[j]
	})

	if len(merged_arr)%2 == 1 {
		return float64(merged_arr[len(merged_arr)/2])
	} else {
		return float64(merged_arr[len(merged_arr)/2]+merged_arr[len(merged_arr)/2-1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
