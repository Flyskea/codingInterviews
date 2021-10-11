package main

import "fmt"

type Trie struct {
	children [2]*Trie
}

func buildDict(nums []int) *Trie {
	root := &Trie{}
	var node *Trie
	for _, num := range nums {
		node = root
		for i := 63; i >= 0; i-- {
			bit := (num >> i) & 1
			if node.children[bit] == nil {
				node.children[bit] = &Trie{}
			}
			node = node.children[bit]
		}
	}
	return root
}

func findMaximumXOR(nums []int) int {
	root := buildDict(nums)
	var (
		node     *Trie
		max, xor int
	)
	for _, num := range nums {
		node = root
		for i := 63; i >= 0; i-- {
			bit := (num >> i) & 1
			if node.children[1-bit] != nil {
				xor = xor<<1 + 1
				node = node.children[1-bit]
			} else {
				xor <<= 1
				node = node.children[bit]
			}
		}
		if max < xor {
			max = xor
		}
	}
	return max
}

func main() {
	fmt.Println(findMaximumXOR([]int{1, 2, 3, 4}))
}
