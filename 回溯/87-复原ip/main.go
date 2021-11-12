package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var (
		sLen        = len(s)
		dfs         func(tart int)
		subRes, ans []string
	)
	if sLen < 4 || sLen > 16 {
		return nil
	}
	dfs = func(start int) {
		if len(subRes) == 4 && start == sLen {
			ans = append(ans, strings.Join(subRes, "."))
			return
		}
		if len(subRes) == 4 && start < sLen {
			return
		}
		for length := 1; length <= 3; length++ {
			if start+length-1 >= sLen {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(start + length)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
