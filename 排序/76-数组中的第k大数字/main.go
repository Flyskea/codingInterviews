package main

import "fmt"

// 快速选择
func quickSelect(nums []int, k, left, right int) int {
	if left >= right {
		return nums[left]
	}
	var (
		i     = left + 1
		j     = right
		pivot = nums[left]
	)
	for {
		for i < right && nums[i] <= pivot {
			i++
		}
		for j > left && nums[j] >= pivot {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	nums[left], nums[j] = nums[j], nums[left]
	if k == j {
		return nums[j]
	}
	if k < j {
		return quickSelect(nums, k, left, j-1)
	}
	return quickSelect(nums, k, j+1, right)
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 9, 11, 2}, 5))
}
