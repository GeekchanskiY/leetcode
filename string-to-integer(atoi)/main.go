package main

import (
	"fmt"
	"math"
	"strings"
)

var chToNum = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'-': 10,
	'+': 11,
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	var (
		brk, rev, ok bool
		res, num     int
	)

	for _, r := range []rune(s) {
		if num, ok = chToNum[r]; ok {
			if num == 10 {
				if brk {
					break
				}

				rev, brk = true, true
			} else if num == 11 {
				if brk {
					break
				}

				brk = true
			} else {
				if res > 214748364 || (res == 214748364 && num > 7) {
					if rev {
						return math.MinInt32
					}
					return math.MaxInt32
				}

				res = res*10 + num

				brk = true
			}
		} else {
			break
		}
	}

	if rev {
		if -res < math.MinInt32 {
			return math.MinInt32
		}

		return -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func main() {
	fmt.Println(myAtoi("42"))
}
