package main

import (
	"container/heap"
	"fmt"
)

type minHeap struct {
	nums1, nums2 []int
	pairs        [][]int
}

func (h *minHeap) Len() int {
	return len(h.pairs)
}

func (h *minHeap) Less(i, j int) bool {
	return h.nums1[h.pairs[i][0]]+h.nums2[h.pairs[i][1]] < h.nums1[h.pairs[j][0]]+h.nums2[h.pairs[j][1]]
}

func (h *minHeap) Swap(i, j int) {
	h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i]
}

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.pairs = append(h.pairs, x.([]int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old.pairs[n-1]
	h.pairs = old.pairs[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 建立最小堆
// 将堆顶元素加入结果
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	minHeap := &minHeap{
		nums1: nums1,
		nums2: nums2,
		pairs: make([][]int, 0),
	}
	// 堆上存索引
	heap.Init(minHeap)
	if len(nums2) > 0 {
		// 升序排列，先试nums1里最小的加上nums2其他的
		for i := 0; i < min(k, len(nums1)); i++ {
			minHeap.Push([]int{i, 0})
		}
	}
	ans := make([][]int, 0)
	// 只需前k个
	for k > 0 && minHeap.Len() != 0 {
		// 弹出最小对
		ids := heap.Pop(minHeap).([]int)
		ans = append(ans, []int{nums1[ids[0]], nums2[ids[1]]})
		// 索引没道nums2的边界
		if ids[1] < len(nums2)-1 {
			heap.Push(minHeap, []int{ids[0], ids[1] + 1})
		}
		k--
	}
	return ans
}

func main() {
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
}
