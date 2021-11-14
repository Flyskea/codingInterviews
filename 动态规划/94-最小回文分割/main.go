package main

import "fmt"

func minCut(s string) int {
	var (
		sLen  = len(s)
		dp    = make([]int, sLen)
		isPal = make([][]bool, sLen)
	)
	// 判断i:j的字串是否为回文
	for i := range isPal {
		isPal[i] = make([]bool, sLen)
		for j := range isPal[i] {
			isPal[i][j] = true
		}
	}
	for i := sLen - 1; i >= 0; i-- {
		for j := i + 1; j < sLen; j++ {
			isPal[i][j] = s[i] == s[j] && isPal[i+1][j-1]
		}
	}
	for i := range dp {
		// 字串是回文跳过
		if isPal[0][i] {
			continue
		}
		dp[i] = i
		for j := 0; j < i; j++ {
			if isPal[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[sLen-1]
}

func main() {
	fmt.Println(minCut("aab"))
}
