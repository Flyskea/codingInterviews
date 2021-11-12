package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var (
		res []string
		dfs func(lindex, rindex int, cur string)
	)
	dfs = func(lindex, rindex int, cur string) {
		if lindex == 0 && rindex == 0 {
			res = append(res, cur)
			return
		}
		if lindex > 0 {
			dfs(lindex-1, rindex, cur+"(")
		}
		if rindex > 0 && lindex < rindex {
			dfs(lindex, rindex-1, cur+")")
		}
	}
	dfs(n, n, "")
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
