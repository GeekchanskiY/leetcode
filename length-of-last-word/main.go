// https://leetcode.com/problems/length-of-last-word

package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	lastLength := 0
	currLength := 0

	for _, ch := range s {
		if ch == ' ' {
			currLength = 0
		} else {
			currLength++
			lastLength = currLength
		}
	}

	return lastLength
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" a sdfg    Helloasdf World  "))
}
