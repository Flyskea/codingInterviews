package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* func binarySearch(arr []int, target, end int) int {
	var (
		low  = 0
		high = end - 1
		mid  int
	)
	for low <= high {
		mid = low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 二分查找
func lenLongestFibSubseq(arr []int) int {
	var (
		arrLen     = len(arr)
		dp         = make([][]int, arrLen)
		ans    int = 2
		prev   int
	)
	for i := range dp {
		dp[i] = make([]int, arrLen)
	}
	for i := 1; i < arrLen; i++ {
		for j := i - 1; j >= 0; j-- {
			prev = arr[i] - arr[j]
			if prev > arr[j] {
				break
			}
			if k := binarySearch(arr, prev, j); k != -1 {
				dp[i][j] = dp[j][k] + 1
			} else {
				dp[i][j] = 2
			}
			ans = max(ans, dp[i][j])
		}
	}
	if ans > 2 {
		return ans
	}
	return 0
} */

// 哈希表
// dp[i][j]表示以i为结尾j为倒数第二个数
// 即a[i]=a[j]-a[k]
func lenLongestFibSubseq(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	res := 0
	hash := make(map[int]int)
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		hash[arr[i]] = i
		dp[i] = make([]int, len(arr))
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			prev := arr[i] - arr[j]
			if prev >= arr[j] {
				break
			}
			if value, ok := hash[prev]; ok {
				dp[i][j] = max(2, dp[j][value]) + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(lenLongestFibSubseq([]int{1, 3, 4, 5, 6, 7, 8}))
}
