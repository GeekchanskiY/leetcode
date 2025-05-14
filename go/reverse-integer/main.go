package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var (
		result = 0
		div    = 10
		divRes int
	)

	for {
		divRes = x % div

		result = result*10 + divRes/(div/10)
		div *= 10

		if divRes == x {
			break
		}
	}

	if result > (math.MaxInt32) || result < math.MinInt32 {
		return 0
	}

	return result
}

func main() {
	fmt.Println(reverse(123456))
	fmt.Println(reverse(-123456))
	fmt.Println(reverse(-12345612312312))
}
