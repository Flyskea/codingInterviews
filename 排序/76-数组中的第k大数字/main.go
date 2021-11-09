package main

import "fmt"

/* // 快速选择
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
} */

func up(nums []int, i int) {
	var (
		parent = (i - 1) >> 1
		tmp    = nums[i]
	)
	for ; parent >= 0 && nums[parent] > tmp; parent = (i - 1) >> 1 {
		nums[i] = nums[parent]
		i = parent
	}
	nums[i] = tmp
}

func down(nums []int, i, N int) {
	var (
		tmp   = nums[i]
		child = i*2 + 1
	)
	for ; i*2+1 <= N; i = child {
		child = i*2 + 1
		if child != N && nums[child+1] < nums[child] {
			child++
		}
		if tmp < nums[child] {
			break
		}
		nums[i] = nums[child]
	}
	nums[i] = tmp
}

func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		up(nums, i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			down(nums, 0, k-1)
		}
	}
	return nums[0]
}

func main() {
	a := []int{2, 3, 1, 4, 0}
	down(a, 0, 4)
	fmt.Println(findKthLargest([]int{3, 2, 1}, 3))
}
