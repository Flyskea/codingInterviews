package main

import "fmt"

// dp[i][j]表示s的前i个字符和t的前j个字符的子序列数目
func numDistinct(s string, t string) int {
	var (
		sLen = len(s)
		tLen = len(t)
		dp   = make([][]int, sLen+1)
	)
	for i := range dp {
		dp[i] = make([]int, tLen+1)
		dp[i][0] = 1
	}
	for i, v1 := range s {
		for j, v2 := range t {
			if v1 == v2 {
				// 如果s[i]==t[j]，则可以匹配s[i]或者不匹配
				dp[i+1][j+1] = dp[i][j+1] + dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	return dp[sLen][tLen]
}

func main() {
	fmt.Println(numDistinct("bagbbag", "bag"))
}
