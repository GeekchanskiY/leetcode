// https://leetcode.com/problems/plus-one

package main

import "fmt"

// this case will not work due to integer overflow...
//func plusOne(digits []int) []int {
//	resultingDigit := int64(0)
//	mul := int64(1)
//
//	for i := len(digits) - 1; i >= 0; i-- {
//		resultingDigit = resultingDigit + int64(digits[i])*mul
//		mul *= 10
//	}
//
//	resultingDigit += 1
//	res := make([]int, 0, len(digits))
//
//	i := 0
//	mul = 10
//	for {
//		res = append(res, int(resultingDigit%mul/(mul/10)))
//		i += 1
//
//		if resultingDigit/mul == 0 {
//			break
//		}
//
//		mul *= 10
//	}
//
//	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
//		res[i], res[j] = res[j], res[i]
//	}
//
//	return res
//}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0

			if i == 0 {
				digits = append(digits, 0)
				digits[0] = 1

				break
			}

			continue
		}

		digits[i] += 1

		break
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
