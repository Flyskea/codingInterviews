package main

import "fmt"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var (
		max int = func() int {
			var tmp int
			for _, v := range arr1 {
				if v > tmp {
					tmp = v
				}
			}
			return tmp
		}()
		ans  = make([]int, 0, len(arr1))
		freq = make([]int, max+1)
	)
	for _, v := range arr1 {
		freq[v]++
	}
	for _, v := range arr2 {
		for ; freq[v] > 0; freq[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range freq {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
