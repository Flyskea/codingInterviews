package main

import (
	"fmt"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	minuteMap := make([]bool, 1440)
	var (
		t                 []string
		hour, minute, tmp int
		minDiff           int = 1439
		prev              int = -1
		first             int = 1439
		last              int = -1
	)
	for _, time := range timePoints {
		t = strings.Split(time, ":")
		hour, _ = strconv.Atoi(t[0])
		minute, _ = strconv.Atoi(t[1])
		tmp = hour*60 + minute
		if minuteMap[tmp] {
			return 0
		}
		minuteMap[tmp] = true
	}
	for i := 0; i < 1440; i++ {
		if minuteMap[i] {
			if prev >= 0 {
				minDiff = min(i-prev, minDiff)
			}
			prev = i
			first = min(first, i)
			last = max(last, i)
		}
	}
	minDiff = min(first+1440-last, minDiff)
	return minDiff
}

func main() {
	fmt.Println(findMinDifference([]string{"01:01", "02:01"}))
}
