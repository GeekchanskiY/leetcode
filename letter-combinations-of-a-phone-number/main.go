package main

import (
	"fmt"
	"strconv"
)

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	charMapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	if len(digits) != 0 {
		result = append(result, "")

		for i := 0; i < len(digits); i++ {
			index, _ := strconv.Atoi(string(digits[i]))
			for len(result[0]) == i {
				permutation := result[0]
				result = result[1:]
				for _, c := range charMapping[index] {
					result = append(result, permutation+string(c))
				}
			}
		}
	}

	return result
}

func main() {
	digits := "237"
	result := letterCombinations(digits)
	fmt.Println(result)
}
