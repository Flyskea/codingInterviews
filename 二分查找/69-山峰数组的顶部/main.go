package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	var (
		mid, tmp  int
		low, high = 1, len(arr) - 2
	)
	for low <= high {
		mid = low + (high-low)>>1
		tmp = arr[mid]
		if tmp > arr[mid+1] && tmp > arr[mid-1] {
			return mid
		} else if tmp < arr[mid-1] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
}
