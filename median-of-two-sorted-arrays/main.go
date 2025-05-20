package main

import (
	"fmt"
)

// // Best solution
//
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	l := len(nums1) + len(nums2)
// 	i1 := 0
// 	i2 := 0
// 	var now int
// 	for i1+i2 < l/2 {
// 		if i1 == len(nums1) {
// 			return findN(nums2, i2, l/2-(i1+i2), l%2)
// 		}
// 		if i2 == len(nums2) {
// 			return findN(nums1, i1, l/2-(i1+i2), l%2)
// 		}
// 		if nums1[i1] < nums2[i2] {
// 			now = nums1[i1]
// 			i1++
// 		} else {
// 			now = nums2[i2]
// 			i2++
// 		}
// 	}
// 	if l%2 != 0 {
// 		if i1 == len(nums1) {
// 			return float64(nums2[i2])
// 		} else if i2 == len(nums2) {
// 			return float64(nums1[i1])
// 		}
// 		if nums1[i1] < nums2[i2] {
// 			return float64(nums1[i1])
// 		} else {
// 			return float64(nums2[i2])
// 		}
// 	}
// 	var next int
// 	if i1 == len(nums1) {
// 		next = nums2[i2]
// 		return float64(now+next) / 2
// 	} else if i2 == len(nums2) {
// 		next = nums1[i1]
// 		return float64(now+next) / 2
// 	}
// 	if nums1[i1] < nums2[i2] {
// 		next = nums1[i1]
// 	} else {
// 		next = nums2[i2]
// 	}
// 	return float64(now+next) / 2
// }
// func findN(nums []int, head int, n int, one int) float64 {
// 	if one == 1 {
// 		return float64(nums[head+n])
// 	} else {
// 		return float64(nums[head+n-1]+nums[head+n]) / 2
// 	}
// }

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
