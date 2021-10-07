package main

import "fmt"

type minHeap struct {
	k    int
	heap []int
}

type KthLargest struct {
	min *minHeap
}

func NewMinHeap(k int, nums []int) *minHeap {
	heap := minHeap{
		k:    k,
		heap: []int{},
	}
	for _, v := range nums {
		heap.add(v)
	}
	return &heap
}

func (mh *minHeap) up(i int) {
	x := mh.heap[i]
	for i > 0 {
		parent := (i - 1) >> 1
		if mh.heap[parent] > x {
			mh.heap[i] = mh.heap[parent]
		} else {
			break
		}
		i = parent
	}
	mh.heap[i] = x
}

func (mh *minHeap) down(i, N int) {
	var tmp, child int
	for tmp = mh.heap[i]; i*2+1 <= N; i = child {
		child = i*2 + 1
		if child != N && mh.heap[child+1] < mh.heap[child] {
			child++
		}
		if tmp < mh.heap[child] {
			break
		}
		mh.heap[i] = mh.heap[child]
	}
	mh.heap[i] = tmp
}

func (mh *minHeap) add(val int) {
	if len(mh.heap) < mh.k {
		mh.heap = append(mh.heap, val)
		mh.up(len(mh.heap) - 1)
	} else if val > mh.heap[0] {
		mh.heap[0] = val
		mh.down(0, mh.k-1)
	}
}

func Constructor(k int, nums []int) KthLargest {
	heap := NewMinHeap(k, nums)
	return KthLargest{
		min: heap,
	}
}

func (kthL *KthLargest) Add(val int) int {
	kthL.min.add(val)
	return kthL.min.heap[0]
}

func main() {
	kthLargest := Constructor(2, []int{0})
	fmt.Println(kthLargest.Add(-1)) // return 4
	fmt.Println(kthLargest.Add(1))  // return 5
	fmt.Println(kthLargest.Add(-2)) // return 5
	fmt.Println(kthLargest.Add(-4)) // return 8
	fmt.Println(kthLargest.Add(3))  // return 8
}
