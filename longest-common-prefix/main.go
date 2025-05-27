package main

import (
	"fmt"
)

// speed O(n * m) / mem O(1)
func longestCommonPrefix(strs []string) string {
	var (
		ethalon   string
		result    []rune
		lowestLen = len(strs[0])
	)

	if len(strs) == 1 {
		return strs[0]
	}

	ethalon = strs[0]
	strs = strs[1:]

	for _, str := range strs {
		if len(str) < lowestLen {
			lowestLen = len(str)
		}
	}

	for i := 0; i < lowestLen; i++ {
		for _, str := range strs {
			if str[i] != ethalon[i] {
				return string(result)
			}
		}

		result = append(result, rune(ethalon[i]))
	}

	return string(result)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
