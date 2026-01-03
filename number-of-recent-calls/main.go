// https://leetcode.com/problems/number-of-recent-calls

package main

import "fmt"

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: make([]int, 0),
	}
}

func (c *RecentCounter) Ping(t int) int {
	c.pings = append(c.pings, t)

	delFrom := 0
	for i := 0; i < len(c.pings); i++ {
		delFrom = i

		if c.pings[i] >= t-3000 {
			break
		}
	}

	c.pings = c.pings[delFrom:]

	return len(c.pings)
}

func main() {
	cnt := Constructor()
	fmt.Println(cnt.Ping(0))
	fmt.Println(cnt.Ping(1000))
	fmt.Println(cnt.Ping(2000))
	fmt.Println(cnt.Ping(3000))
	fmt.Println(cnt.Ping(4000))
	fmt.Println(cnt.Ping(10000))
}
