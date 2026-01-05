// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string

package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	p1 := 0
	p2 := 0
	tmpP := 0

	for p1 < len(haystack) {
		tmpP = p1
		p2 = 0

		for p2 < len(needle) {
			if len(haystack) <= tmpP {
				return -1
			}

			if haystack[tmpP] != needle[p2] {
				break
			}

			p2++
			tmpP++
		}

		if p2 == len(needle) {
			return p1
		}

		p1++
	}

	return -1
}

func strStrCheat(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("hello world", "hello"))
	fmt.Println(strStr("adgasdfgasdfasdfhelloworld", "hello"))
	fmt.Println(strStr("hello world", "GeekchanskiY"))
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("aaa", "aaaa"))
}
