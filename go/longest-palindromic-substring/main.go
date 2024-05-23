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
	longest_found := ""
	for i := 0; i < len(s); i++ {
		// Check for odd length palindromes
		for current_recursion_level := 0; ; current_recursion_level++ {
			if i-current_recursion_level >= 0 && i+current_recursion_level < len(s) {
				if isPalindrome(s[i-current_recursion_level : i+current_recursion_level+1]) {
					if len(longest_found) < 2*current_recursion_level+1 {
						longest_found = s[i-current_recursion_level : i+current_recursion_level+1]
					}
				} else {
					break
				}
			} else {
				break
			}
		}
		// Check for even length palindromes
		for current_recursion_level := 0; ; current_recursion_level++ {
			if i-current_recursion_level >= 0 && i+current_recursion_level+1 < len(s) {
				if isPalindrome(s[i-current_recursion_level : i+current_recursion_level+2]) {
					if len(longest_found) < 2*current_recursion_level+2 {
						longest_found = s[i-current_recursion_level : i+current_recursion_level+2]
					}
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	return longest_found
}

func main() {
	fmt.Println(longestPalindrome("A man, a plan, a canal: Panama"))
}
