// https://leetcode.com/problems/merge-strings-alternately

package main

import "fmt"

// mergeAlternately
func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}

	if len(word2) == 0 {
		return word1
	}

	curr, i1, i2 := 0, 0, 0

	// strings.Builder is slightly faster here
	res := make([]uint8, len(word1)+len(word2))

	for i1 < len(word1) || i2 < len(word2) {
		if i1 >= len(word1) {
			res[curr] = word2[i2]
			i2++
			curr++

			continue
		}

		if i2 >= len(word2) {
			res[curr] = word1[i1]
			i1++
			curr++

			continue
		}

		if curr%2 != 1 {
			res[curr] = word1[i1]
			i1++
		} else {
			res[curr] = word2[i2]
			i2++
		}

		curr++
	}

	return string(res)
}

func main() {
	fmt.Println(mergeAlternately("hello", "world"))
}
