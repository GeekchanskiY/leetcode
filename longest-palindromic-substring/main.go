package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	isPal := true

	if len(s)%2 == 1 {
		for i := 0; i < (len(s)-1)/2; i++ {
			if s[len(s)-i-1] != s[i] {
				isPal = false
				break
			}
		}
	} else {
		for i := 0; i < len(s)/2; i++ {
			if s[len(s)-i-1] != s[i] {
				isPal = false
				break
			}
		}
	}
	return isPal
}

func longestPalindrome(s string) string {
	longestFound := ""

	for i := 0; i < len(s); i++ {
		// Check for odd length palindromes
		for currentRecursionLevel := 0; ; currentRecursionLevel++ {
			if i-currentRecursionLevel >= 0 && i+currentRecursionLevel < len(s) {
				if isPalindrome(s[i-currentRecursionLevel : i+currentRecursionLevel+1]) {
					if len(longestFound) < 2*currentRecursionLevel+1 {
						longestFound = s[i-currentRecursionLevel : i+currentRecursionLevel+1]
					}
				} else {
					break
				}
			} else {
				break
			}
		}

		// Check for even length palindromes
		for currentRecursionLevel := 0; ; currentRecursionLevel++ {
			if i-currentRecursionLevel >= 0 && i+currentRecursionLevel+1 < len(s) {
				if isPalindrome(s[i-currentRecursionLevel : i+currentRecursionLevel+2]) {
					if len(longestFound) < 2*currentRecursionLevel+2 {
						longestFound = s[i-currentRecursionLevel : i+currentRecursionLevel+2]
					}
				} else {
					break
				}
			} else {
				break
			}
		}
	}

	return longestFound
}

func main() {
	fmt.Println(longestPalindrome("A man, a plan, a canal: Panama"))
}
