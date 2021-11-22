package main

import (
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	inDegrees := make([]int, numCourses)
	queue := make([]int, 0)
	ans := make([]int, 0)
	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegrees[v[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		ans = append(ans, course)
		for _, next := range graph[course] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if len(ans) == numCourses {
		return ans
	}
	return nil
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
}
