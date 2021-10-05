package main

import (
	"fmt"
)

func isAlienSorted(words []string, order string) bool {
	orderMap := make([]int, 26)
	for i, v := range order {
		orderMap[v-'a'] = i
	}
	var (
		word1, word2 string
		ch1, ch2     byte
		flag         bool
	)
	wordsLen := len(words)
	for i := 0; i < wordsLen-1; i++ {
		word1 = words[i]
		word2 = words[i+1]
		flag = false
		for j := 0; j < len(word1) && j < len(word2); j++ {
			ch1 = word1[j]
			ch2 = word2[j]
			if orderMap[ch1-'a'] < orderMap[ch2-'a'] {
				flag = true
				break
			}
			if orderMap[ch1-'a'] > orderMap[ch2-'a'] {
				return false
			}
		}
		if !flag && len(word1) > len(word2) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))
}
