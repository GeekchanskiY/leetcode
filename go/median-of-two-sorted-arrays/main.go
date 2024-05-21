package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merged_arr []int

	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) == 0 {
			merged_arr = append(merged_arr, nums2[0])
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			merged_arr = append(merged_arr, nums1[0])
			nums1 = nums1[1:]
		} else if nums1[0] <= nums2[0] {
			merged_arr = append(merged_arr, nums1[0])
			nums1 = nums1[1:]
		} else {
			merged_arr = append(merged_arr, nums2[0])
			nums2 = nums2[1:]
		}
	}
	fmt.Println(merged_arr)

	if len(merged_arr)%2 == 1 {
		return float64(merged_arr[len(merged_arr)/2])
	} else {
		return float64(merged_arr[len(merged_arr)/2]+merged_arr[len(merged_arr)/2-1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
