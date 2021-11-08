package main

import (
	"fmt"
	"math"
)

func getHours(piles []int, speed int) int {
	var hours int
	for _, v := range piles {
		hours += (v + speed - 1) / speed
	}
	return hours
}

func minEatingSpeed(piles []int, h int) int {
	var (
		low, high, mid, hours int
	)
	for _, v := range piles {
		if v > high {
			high = v
		}
	}
	for low < high {
		mid = low + (high-low)>>1
		hours = getHours(piles, mid)
		if hours <= h {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(math.Ceil(3.2))
}
