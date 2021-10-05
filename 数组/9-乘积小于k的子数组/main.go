package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	var (
		left, right, count int
		product            = 1
	)
	for ; right < len(nums); right++ {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		if left <= right {
			count += right - left + 1
		}
	}
	return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
