package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i][j]表示text1前i个字符与text2前j个字符的最长子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	var (
		text1Len = len(text1)
		text2Len = len(text2)
		dp       = make([][]int, text1Len+1)
	)
	for i := range dp {
		dp[i] = make([]int, text2Len+1)
	}
	for i := 0; i < text1Len; i++ {
		for j := 0; j < text2Len; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[text1Len][text2Len]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
