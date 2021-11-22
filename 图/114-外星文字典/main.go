package main

import (
	"fmt"
	"strings"
)

func alienOrder(words []string) string {
	graph := make(map[rune]map[rune]bool)
	inDegrees := make(map[rune]int)
	for _, word := range words {
		for _, c := range word {
			if _, has := graph[c]; !has {
				graph[c] = make(map[rune]bool)
				inDegrees[c] = 0
			}
		}
	}
	n := len(words)
	for i := 1; i < n; i++ {
		var j int
		length := min(len(words[i-1]), len(words[i]))
		for ; j < length; j++ {
			ch1 := rune(words[i-1][j])
			ch2 := rune(words[i][j])
			if ch1 != ch2 {
				if _, has := graph[ch1][ch2]; !has {
					graph[ch1][ch2] = true
					inDegrees[ch2]++
				}
				break
			}
		}
		if j == length && len(words[i-1]) > len(words[i]) {
			return ""
		}
	}

	var ret strings.Builder
	var queue []rune
	for i, d := range inDegrees {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		ret.WriteRune(ch)
		for c := range graph[ch] {
			inDegrees[c]--
			if inDegrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	if ret.Len() != len(inDegrees) {
		return ""
	}
	return ret.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
}
