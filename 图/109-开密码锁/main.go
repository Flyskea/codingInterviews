package main

import (
	"fmt"
)

/* func getNeighbors(cur string) (nexts []string) {
	curByte := []byte(cur)
	tmp := []byte{}
	tmp = append(tmp, curByte...)
	for i := range curByte {
		if curByte[i] == '0' {
			curByte[i] = '9'
		} else {
			curByte[i] -= 1
		}
		nexts = append(nexts, string(curByte))
		copy(curByte, tmp)
		if curByte[i] == '9' {
			curByte[i] = '0'
		} else {
			curByte[i] += 1
		}
		nexts = append(nexts, string(curByte))
		copy(curByte, tmp)
	}
	return
}

func openLock(deadends []string, target string) int {
	var (
		deadMap = make(map[string]bool)
		visited = make(map[string]bool)
		init    = "0000"
		step    int
		cur     string
	)
	for _, v := range deadends {
		deadMap[v] = true
	}
	if deadMap[init] || deadMap[target] {
		return -1
	}
	queue := []string{}
	queue = append(queue, init)
	visited[init] = true
	for len(queue) != 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			cur = queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			nexts := getNeighbors(cur)
			for _, v := range nexts {
				if !deadMap[v] && !visited[v] {
					queue = append(queue, v)
					visited[v] = true
				}
			}
		}
		step++
	}
	return -1
} */

// 双向bfs
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	visited := make(map[string]bool)
	deads := make(map[string]bool)
	step := 0
	visited["0000"] = true
	q1["0000"] = true
	q2[target] = true

	for _, v := range deadends {
		deads[v] = true
	}

	for len(q1) > 0 && len(q2) > 0 {
		temp := make(map[string]bool)
		for v := range q1 {
			if q2[v] {
				return step
			}
			if deads[v] {
				continue
			}

			visited[v] = true

			for i := 0; i < 4; i++ {
				plus := plusOne(v, i)
				if !visited[plus] {
					temp[plus] = true

				}

				sub := subOne(v, i)
				if !visited[sub] {
					temp[sub] = true
				}
			}

		}

		q1 = q2
		q2 = temp
		step++

	}

	return -1
}

func plusOne(curr string, idx int) string {
	tmp := []byte(curr)
	b := tmp[idx]
	b++

	if b > '9' {
		b = '0'
	}

	tmp[idx] = b
	return string(tmp)
}

func subOne(curr string, idx int) string {
	tmp := []byte(curr)
	b := tmp[idx]
	b--

	if b < '0' {
		b = '9'
	}

	tmp[idx] = b
	return string(tmp)
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
