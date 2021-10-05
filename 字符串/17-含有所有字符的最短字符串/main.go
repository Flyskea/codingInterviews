package main

import "fmt"

func minWindow(s string, t string) string {
	var (
		freq                   = make([]int, 128)
		flag                   = make([]bool, 128)
		cnt, left, right, minL int
		sLen                   = len(s)
		minSize                = sLen + 1
	)
	for i := range t {
		flag[t[i]] = true
		freq[t[i]]++
	}
	minSize = len(s) + 1
	for right < sLen {
		if flag[s[right]] {
			freq[s[right]]--
			if freq[s[right]] >= 0 {
				cnt++
			}
			for cnt == len(t) {
				if right-left+1 < minSize {
					minL = left
					minSize = right - left + 1
				}
				if flag[s[left]] {
					freq[s[left]]++
					if freq[s[left]] > 0 {
						cnt--
					}
				}
				left++
			}
		}
		right++
	}
	if minSize > len(s) {
		return ""
	}
	return s[minL : minL+minSize]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "BANC"))
}
