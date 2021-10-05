package main

import "fmt"

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}
	return start >= end
}

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	mid := len(s) / 2
	for start < mid {
		if s[start] != s[end] {
			break
		}
		start++
		end--
	}
	return start == mid || isPalindrome(s, start+1, end) || isPalindrome(s, start, end-1)
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
