package main

import "fmt"

// 设 f(i, j) 表示字符串 s1 的下标从 0 到 i 的子字符串 s1[0] ~ s1[i] （长度为 i + 1)
// 和字符串 s2 的下标从 0 到 j 的子字符串 s2[0] ~ s2[j] （长度为 j + 1)
// 能否交织成字符串 s3 的下标从 0 到 i + j + 1 的子字符串 s3[0] ~ s3[i + j + 1] （长度为 i + j + 2)
func isInterleave(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if (s1Len + s2Len) != s3Len {
		return false
	}
	dp := make([][]bool, s1Len+1)
	for i := 0; i <= s1Len; i++ {
		dp[i] = make([]bool, s2Len+1)
	}
	dp[0][0] = true
	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[s1Len][s2Len]
}

func main() {
	fmt.Println(isInterleave("a", "abc", "aabc"))
}
