package main

import "math/rand"

type Solution struct {
	sums  []int
	total int
}

func Constructor(w []int) Solution {
	total := 0
	sums := make([]int, len(w))
	for i, v := range w {
		total += v
		sums[i] = total
	}
	return Solution{
		sums:  sums,
		total: total,
	}
}

func (s *Solution) PickIndex() int {
	var (
		low, mid int
		high     = len(s.sums)
		p        = rand.Intn(s.total)
	)
	for low <= high {
		mid = low + (high-low)>>1
		if s.sums[mid] > p {
			if mid == 0 || (s.sums[mid-1] <= p) {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
