package main

import (
	"fmt"
)

// isMatch
// s - only lowercase English letters
// p - only lowercase English letters, and '.' or '*'
// It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
// It must match the entire string! (makes it easier)
func isMatch(s string, p string) bool {
	var (
		i                 int
		strIdx, newStrIdx int
		currentRegex      rune
	)

	if p == s {
		return true
	}

	for i = 0; i < len(p); i++ {
		if strIdx >= len(s) {
			return false
		}

		currentRegex = rune(p[i])

		if i+1 < len(p) && p[i+1] == '*' {
			newStrIdx = processStar(s, strIdx, rune(p[i]))

			// process if .* has the next character
			if newStrIdx == len(s) {
				if len(p) == i+2 {
					return true
				}

				i++
				strIdx += 1

				continue
			}

			strIdx = newStrIdx

			i++

			continue
		}

		if currentRegex == '.' {
			strIdx += 1

			continue
		}

		if rune(s[strIdx]) != currentRegex {
			return false
		}

		strIdx += 1
	}

	if strIdx != len(s) {
		return false
	}

	return true
}

// processStar returns if * matches, and the next position to process
func processStar(s string, pos int, starVal rune) int {
	if starVal == '.' {
		return len(s)
	}

	for i := pos; i < len(s); i++ {
		if rune(s[i]) != starVal {
			return i
		}
	}

	return len(s)
}

func restoreDotStar(s string, pos int) {

}

func main() {
	// fmt.Println(isMatch("aa", "a"))
	// fmt.Println(isMatch("aa", "a*"))
	// fmt.Println(isMatch("aab", "a*"))
	fmt.Println(isMatch("aaa", "a*a"))
	// fmt.Println(isMatch("aabbcc", "aabbcc"))
	// fmt.Println(isMatch("ab", ".*"))

	// fmt.Println(isMatch("mississippi", "mis*is*p*."))
	//
	// fmt.Println(isMatch("aab", "c*a*b"))

	// fmt.Println(isMatch("ab", ".*c"))
	// fmt.Println(isMatch("cab", "c.*"))
}
