package main

import "fmt"

func similar(str1, str2 string) bool {
	diff := 0
	for i := range str1 {
		if str1[i] != str2[i] {
			diff++
		}
	}
	return diff <= 2
}

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

func Union(disjointSet []int, root1, root2 int) {
	if disjointSet[root1] > disjointSet[root2] {
		disjointSet[root2] += disjointSet[root1]
		disjointSet[root1] = root2
	} else {
		disjointSet[root1] += disjointSet[root2]
		disjointSet[root2] = root1
	}
}

func numSimilarGroups(strs []string) int {
	var (
		ans         int
		strsLen     = len(strs)
		disjointSet = make([]int, strsLen)
	)
	for i := range disjointSet {
		disjointSet[i] = -1
	}
	for i := 0; i < strsLen; i++ {
		for j := i + 1; j < strsLen; j++ {
			if similar(strs[i], strs[j]) {
				node1, node2 := Find(disjointSet, i), Find(disjointSet, j)
				if node1 != node2 {
					Union(disjointSet, node1, node2)
				}
			}
		}
	}
	for _, v := range disjointSet {
		if v < 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}
