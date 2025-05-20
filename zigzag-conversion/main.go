package main

import (
	"strings"
)

// https://docs.google.com/spreadsheets/d/1KTbVslo-2oXSWEpUGoSCyXUjV8N4IXhQUR9ACiLIjS0/edit?usp=sharing

// I almost found mathemathical dependency for this problem, but I got really tired on finding the dependencies....
func convert_old(s string, numRows int) string {
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
		if even == true {

			rf = numRows*2 - (4 * j) + 1
			for i := 0; i < len(s)-1; i++ {
				if counter == 0 {
					sb.WriteString(string(s[i]))
				}
				counter += 1
				if counter > rf {
					counter = 0
				}
			}
			even = false
		} else {
			rf = numRows*2 - (4 * j) + 2
			for i := 0; i < len(s)-1; i++ {
				if counter == 0 {
					sb.WriteString(string(s[i]))
				}
				counter += 1
				if counter > rf {
					counter = 0
				}
			}
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

// Submitted solution
func convert(s string, numRows int) string {
	if numRows < 2 || len(s) < 2 {
		return s
	}

	rows := make([]string, numRows)

	goingDown := false
	row := 0

	for _, char := range s {
		rows[row] += string(char)

		if (row == 0) || (row == numRows-1) {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {

	println(convert("PAYPALISHIRING", 3))
}
