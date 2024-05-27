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

func main() {
	fmt.Println(isPalindrome("bbab"))
}
