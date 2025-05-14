package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mergedArr []int

	for len(nums1) > 0 || len(nums2) > 0 {
		switch {
		case len(nums1) == 0:
			mergedArr = append(mergedArr, nums2[0])
			nums2 = nums2[1:]
		case len(nums2) == 0:
			mergedArr = append(mergedArr, nums1[0])
			nums1 = nums1[1:]
		case nums1[0] <= nums2[0]:
			mergedArr = append(mergedArr, nums1[0])
			nums1 = nums1[1:]
		default:
			mergedArr = append(mergedArr, nums2[0])
			nums2 = nums2[1:]
		}
	}

	if len(mergedArr)%2 == 1 {
		return float64(mergedArr[len(mergedArr)/2])
	} else {
		return float64(mergedArr[len(mergedArr)/2]+mergedArr[len(mergedArr)/2-1]) / 2
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
