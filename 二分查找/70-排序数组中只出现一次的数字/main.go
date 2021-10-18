package main

import (
	"fmt"
)

/* // 含有只出现一个数字的数组一定是奇数
func singleNonDuplicate(nums []int) int {
	var (
		low, mid, tmp int
		high          = len(nums) - 1
		halvesAreEven bool
	)
	for low < high {
		mid = low + (high-low)>>1
		// 单个元素是在mid左边还是右边
		halvesAreEven = (high-mid)%2 == 0
		tmp = nums[mid]
		if nums[mid+1] == tmp {
			// 左边子数组为偶数，low更新为右边子数组开始位置
			if halvesAreEven {
				low = mid + 2
				// 否则更新high为左边子数组开始位置
			} else {
				high = mid - 1
			}
		} else if nums[mid-1] == tmp {
			if halvesAreEven {
				high = mid - 2
			} else {
				low = mid + 1
			}
		} else {
			return tmp
		}
	}
	return nums[low]
} */

// 奇数长度的数组首尾元素索引都为偶数
func singleNonDuplicate(nums []int) int {
	var (
		low, mid int
		high     = len(nums) - 1
	)
	for low < high {
		mid = low + (high-low)>>1
		if mid%2 == 1 {
			mid -= 1
		}
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 2, 2, 3, 3}))
}
