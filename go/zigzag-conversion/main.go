package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	var sb strings.Builder
	if numRows == 1 {
		return s
	}
	iterations := numRows*2 - 2
	var counter int = 0

	// First line
	for i := 0; i < len(s); i++ {
		if counter == 0 {
			sb.WriteString(string(s[i]))
		}
		counter += 1
		if counter >= iterations {
			counter = 0
		}
	}

	// Middle lines (skips if numRows == 2, SWAG)

	even := true
	var rf int // Reducion factor
	for j := 1; j < numRows-1; j++ {
		counter = -j
		if even {

			rf = numRows*2 - (4 * j)
			for i := 0; i < len(s); i++ {
				if counter == 0 {
					sb.WriteString(string(s[i]))
					println("a")
				}
				counter += 1
				if counter > rf {
					counter = 0
				}
			}
			even = false
		} else {
			even = true
		}
	}

	// Last line
	counter = -numRows + 1
	for i := 0; i < len(s); i++ {
		if counter == 0 {
			sb.WriteString(string(s[i]))
		}
		counter += 1
		if counter >= iterations {
			counter = 0
		}
	}

	return sb.String()
}

func main() {

	println(convert("PAYPALISHIRING", 3))
}
