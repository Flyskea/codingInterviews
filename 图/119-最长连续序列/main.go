package main

import "fmt"

func Find(disjointSet []int, x int) int {
	root := x
	tmp := 0
	for disjointSet[root] >= 0 {
		root = disjointSet[root]
	}
	for x != root {
		tmp = disjointSet[x]
		disjointSet[x] = root
		x = tmp
	}
	return root
}

func Union(disjointSet []int, x1, x2 int) {
	root1, root2 := Find(disjointSet, x1), Find(disjointSet, x2)
	if disjointSet[root1] > disjointSet[root2] {
		disjointSet[root2] += disjointSet[root1]
		disjointSet[root1] = root2
	} else {
		disjointSet[root1] += disjointSet[root2]
		disjointSet[root2] = root1
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		disjointSet = make([]int, len(nums))
		numsMap     = make(map[int]int)
		ans, id     int
	)
	for i := range disjointSet {
		disjointSet[i] = -1
	}
	for _, v := range nums {
		if _, ok := numsMap[v]; !ok {
			numsMap[v] = id
			if id1, ok := numsMap[v-1]; ok {
				Union(disjointSet, id, id1)
			}
			if id2, ok := numsMap[v+1]; ok {
				Union(disjointSet, id, id2)
			}
			id++
		}
	}
	for _, v := range disjointSet {
		if -v > ans {
			ans = -v
		}
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 101, 4, 200, 1, 3, 2, 5}))
}
