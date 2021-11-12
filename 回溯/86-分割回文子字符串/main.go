package main

import (
	"fmt"
)

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string) [][]string {
	var (
		dfs  func(start int, cur []string)
		ans  [][]string
		sLen = len(s)
	)
	dfs = func(start int, cur []string) {
		if start == sLen {
			ans = append(ans, append([]string{}, cur...))
		}
		for i := start; i < sLen; i++ {
			if isPalindrome(s, start, i) {
				cur = append(cur, s[start:i+1])
				dfs(i+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0, []string{})
	return ans
}

func main() {
	fmt.Println(partition("google"))
}
