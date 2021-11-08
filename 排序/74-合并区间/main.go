package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var (
		i, j, end    int
		intervalsLen = len(intervals)
		ans          [][]int
	)
	if intervalsLen < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	for i < intervalsLen {
		end = intervals[i][1]
		for j = i + 1; j < intervalsLen && intervals[j][0] <= end; j++ {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
		}
		ans = append(ans, []int{intervals[i][0], end})
		i = j
	}
	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
