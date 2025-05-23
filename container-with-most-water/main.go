// Space Complexity: O(1)
// Time Complexity: O(n)
package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var p1, p2, area, minHeight, currentArea = 0, len(height) - 1, 0, 0, 0

	for p1 < len(height)-1 && p2 > 0 {
		minHeight = min(height[p1], height[p2])

		currentArea = minHeight * (p2 - p1)

		if currentArea > area {
			area = currentArea
		}

		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return area
}

func maxArea2(height []int) int {
	var res, tmp, i, j = 0, 0, 0, len(height) - 1

	for i <= j {

		tmp = min(height[i], height[j]) * (j - i)

		if tmp > res {
			res = tmp
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}

	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 7 * 7 = 49
	fmt.Println(maxArea([]int{3, 6, 1}))                   // 1 * 3 = 3
	fmt.Println(maxArea([]int{8, 7, 2, 1}))                // 7 * 1 = 7
}
