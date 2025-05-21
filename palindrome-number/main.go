package main

import (
	"fmt"
	"math"
)

func isPalindromeFast(x int) bool {
	if x < 0 {
		return false // Negative numbers are not palindromes
	}

	original := x
	reversed := 0

	// Reverse the digits of the number
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// If the reversed number is equal to the original, it's a palindrome
	return original == reversed
}

// It was a genius, but did not work :(
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	// index2 counts int len
	index1, index2 := 1, int(math.Pow(10, math.Ceil(math.Log10(float64(x)))))
	if index2 == 100 {
		return x%10 == x/10
	}

	num1, num2 := 0, 0
	kf := 1
	isEvenLen := index2%2 == 0

	for {
		index1 *= 10
		index2 /= 10
		if isEvenLen {
			if index1*10 == index2 {
				num1 = x / index2
				num2 = x % index1
				if kf != 1 {
					num1 %= 10
					num2 /= kf

				}

				if num1 != num2 {
					return false
				}

				return true
			}
		} else {
			if index2 == index1 {
				return true
			}
		}

		num1 = x / index2
		num2 = x % index1
		if kf != 1 {
			if num1 != 0 {
				num1 %= 10
			}
			if num2 != 0 {
				num2 /= kf
			}
		}

		if num1 != num2 {
			return false
		}

		kf *= 10

	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1001))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(11))
	fmt.Println(isPalindrome(121))

	fmt.Println(isPalindromeFast(10))
	fmt.Println(isPalindromeFast(121))
}
