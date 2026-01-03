// https://leetcode.com/problems/is-subsequence

package main

import "fmt"

func isSubsequence(s string, t string) bool {
	p1 := 0
	p2 := 0
	for p1 < len(s) && p2 < len(t) {
		if s[p1] == t[p2] {

			p1++
			p2++

			continue
		}

		p2++
	}

	if p1 == len(s) {
		return true
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "anmbkjgc"))
}
