package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mergedArr []int

	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) == 0 {
			mergedArr = append(mergedArr, nums2[0])
			nums2 = nums2[1:]
		} else if len(nums2) == 0 {
			mergedArr = append(mergedArr, nums1[0])
			nums1 = nums1[1:]
		} else if nums1[0] <= nums2[0] {
			mergedArr = append(mergedArr, nums1[0])
			nums1 = nums1[1:]
		} else {
			mergedArr = append(mergedArr, nums2[0])
			nums2 = nums2[1:]
		}
	}
	fmt.Println(mergedArr)

	if len(mergedArr)%2 == 1 {
		return float64(mergedArr[len(mergedArr)/2])
	} else {
		return float64(mergedArr[len(mergedArr)/2]+mergedArr[len(mergedArr)/2-1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
