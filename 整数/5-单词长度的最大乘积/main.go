package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(words []string) int {
	flag := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			flag[i] |= 1 << (ch - 'a')
		}
	}
	result := 0
	wordsLen := len(words)
	for i := 0; i < wordsLen; i++ {
		for j := i + 1; j < wordsLen; j++ {
			if flag[i]&flag[j] == 0 {
				result = max(result, len(words[i])*len(words[j]))
			}
		}
	}
	return result
}

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}))
}
