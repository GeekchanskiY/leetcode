package main

import (
	"fmt"
)

var symbolValue = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var (
		l                          = len(s)
		r                          rune
		recentVal, currentVal, res int
	)

	for i := 0; i < l; i++ {
		r = rune(s[i])

		currentVal = symbolValue[r]

		if currentVal > recentVal && recentVal != 0 {
			res += currentVal - recentVal

			currentVal = 0
		} else {
			res += recentVal

			recentVal = currentVal
		}

	}

	if recentVal != 0 {
		res += recentVal
	}

	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("XIV"))
	fmt.Println(romanToInt("XXXIII"))
	fmt.Println(romanToInt("XXXVIII"))
}
