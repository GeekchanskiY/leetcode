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
		m, n = len(s), len(p)
		dp   = make([][]bool, m+1)
	)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	// Handle cases where pattern starts with '*' (only valid when preceding character exists)
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // '*' can match empty sequence
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' { // Direct match or wildcard '.'
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // Handle '*' properly
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return dp[m][n]
}

func main() {
	// fmt.Println(isMatch("aa", "a"))
	// fmt.Println(isMatch("aa", "a*"))
	// fmt.Println(isMatch("aab", "a*"))
	// fmt.Println(isMatch("aaa", "a*a"))
	// fmt.Println(isMatch("aabbcc", "aabbcc"))
	// fmt.Println(isMatch("ab", ".*"))

	// fmt.Println(isMatch("mississippi", "mis*is*p*."))

	// fmt.Println(isMatch("aab", "c*a*b"))

	// fmt.Println(isMatch("ab", ".*c"))
	// fmt.Println(isMatch("cab", "c.*"))
	// fmt.Println(isMatch("aab", "c***a*b") == true)
	// fmt.Println(isMatch("a", "ab*") == true)
	// fmt.Println(isMatch("a", "ab*a") == false)
	// fmt.Println(isMatch("a", "ab*a*c*") == true)
	// fmt.Println(isMatch("a", ".*..a*") == false)
	// fmt.Println(isMatch("a", "..a*"))
	// fmt.Println(isMatch("abbbcd", "ab*bbbcd"))
	// fmt.Println(isMatch("abcd", "d*"))
	fmt.Println(isMatch("abbabaaaaaaacaa", "a*.*b.a.*c*b*a*c*"))

	fmt.Println(isMatch("ab", ".*c"))
}
