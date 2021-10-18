package main

import "fmt"

func searchInsert(nums []int, target int) int {
	var (
		mid, tmp  int
		low, high = 0, len(nums) - 1
	)
	for low <= high {
		mid = low + (high-low)>>1
		tmp = nums[mid]
		if tmp == target {
			return mid
		} else if tmp > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
