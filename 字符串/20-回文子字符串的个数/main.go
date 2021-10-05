package main

import "fmt"

func countSubstrings(s string) int {
	ans := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := 0; j < 2; j++ {
			left := i
			right := i + j
			for left >= 0 && right < sLen && s[left] == s[right] {
				ans++
				left--
				right++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
