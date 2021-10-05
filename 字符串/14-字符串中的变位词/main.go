package main

import "fmt"

func checkInclusion(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	counts := make([]int, 26)
	for i := range s1 {
		counts[s1[i]-'a']++
		counts[s2[i]-'a']--
	}

	if areAllZero(counts) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		counts[s2[i]-'a']--
		counts[s2[i-len(s1)]-'a']++
		if areAllZero(counts) {
			return true
		}
	}
	return false
}

func areAllZero(nums []int) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkInclusion("bc", "akflascb"))
}
