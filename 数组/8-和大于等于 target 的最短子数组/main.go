package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	var left, right, sum, numsLen, min, tmp int
	numsLen = len(nums)
	min = math.MaxInt64
	for right < numsLen {
		sum += nums[right]
		for left <= right && sum >= target {
			tmp = right - left + 1
			if tmp < min {
				min = tmp
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func main() {
	fmt.Println(minSubArrayLen(2, []int{0, 1, 1, 2, 3, -1}))
}
