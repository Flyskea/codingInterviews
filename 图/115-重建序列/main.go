package main

import "fmt"

func sequenceReconstruction(org []int, seqs [][]int) bool {
	if len(seqs) == 0 {
		return false
	}
	graph := make(map[int]map[int]bool)
	inDegrees := make([]int, len(org)+1)
	var queue []int
	n := len(org)
	for _, seq := range seqs {
		for _, v := range seq {
			if v < 1 || v > n {
				return false
			}
		}
		i := seq[0]
		for _, j := range seq[1:] {
			if _, ok := graph[i][j]; !ok {
				if graph[i] == nil {
					graph[i] = map[int]bool{}
				}
				graph[i][j] = true
				inDegrees[j]++
			}
			i = j
		}
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	cnt := 0
	for len(queue) == 1 {
		if org[cnt] != queue[0] {
			return false
		}
		cur := queue[0]
		queue = queue[1:]
		cnt++
		for next := range graph[cur] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return cnt == n
}

func main() {
	fmt.Println(sequenceReconstruction([]int{1}, [][]int{{2}}))
}
