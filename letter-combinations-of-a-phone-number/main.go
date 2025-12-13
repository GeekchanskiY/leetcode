package main

import "fmt"

var (
	digitsMap = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 1 {
		return digitsMap[digits]
	}

	totalCombinations := len(digitsMap[string(digits[0])])
	for _, digit := range digits[1:] {
		totalCombinations *= len(digitsMap[string(digit)])
	}

	sourceLists := make([]string, 0, totalCombinations)
	firstDigitLen := totalCombinations / len(digitsMap[string(digits[0])])
	for _, digit := range digitsMap[string(digits[0])] {
		for i := 0; i < firstDigitLen; i++ {
			sourceLists = append(sourceLists, digit)
		}
	}

	currentOccurrences := totalCombinations

	for _, src := range digits[1:] {
		iterDigits := digitsMap[string(src)]
		dLen := len(iterDigits)
		currentOccurrences /= dLen

		for i, digit := range iterDigits {
			for z := 0; z < dLen; z++ {
				// TODO: fix this formula
				sourceLists[(i*dLen)+(i*currentOccurrences)+z] = sourceLists[(i*dLen)+(i*currentOccurrences)+z] + digit
			}
		}

	}

	// 0 1 2 . . . . . . 9 10 11 . . . . . . 18 19 20 . . . . . . .
	// . . . 3 4 5 . . . . . . 12 13 14 . . . . . . 21 22 23 . . .
	// . . . . . . 6 7 8 . . . . . . 15 16 17 . . . . . . 24 25 26

	return sourceLists
}

func main() {
	digits := "234"
	result := letterCombinations(digits)
	fmt.Println(result)
}
