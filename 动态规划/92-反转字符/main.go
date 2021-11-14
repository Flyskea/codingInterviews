package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFlipsMonoIncr(s string) int {
	// 记录最小翻转次数 和 1的个数，每次遇到1计数+1，
	// 遇到0考虑是翻转所有的1 还是 翻转这个0
	var dp, cnt int
	for _, v := range s {
		if v == '1' {
			cnt++
		} else {
			dp = min(dp+1, cnt)
		}
	}
	return dp
}

func main() {
	fmt.Println(minFlipsMonoIncr("00110"))
}
