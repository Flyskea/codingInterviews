package main

import (
	"fmt"
)

func mySqrt(x int) int {
	var (
		low      = 1
		high     = x
		mid, ans int
	)
	for low <= high {
		mid = low + (high-low)/2
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(mySqrt(9))
}
