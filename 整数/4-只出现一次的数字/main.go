package main

import "fmt"

func singleNumber(nums []int) int {
	bitSums := make([]int, 64)
	for _, num := range nums {
		for i := 0; i < 64; i++ {
			bitSums[i] += (num >> (63 - i)) & 1
		}
	}
	result := 0
	for i := 0; i < 64; i++ {
		result = (result << 1) + bitSums[i]%3
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 3, 2, 1, 2, 2}))
}
