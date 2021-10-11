package main

import "fmt"

type Trie struct {
	children [26]*Trie
	end      bool
}

func buildDict(dictionary []string) *Trie {
	root := &Trie{}
	var node *Trie
	for _, v := range dictionary {
		node = root
		// 相同后缀，将单词翻转存入
		for i := len(v) - 1; i >= 0; i-- {
			ch := v[i]
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.end = true
	}
	return root
}

func minimumLengthEncoding(words []string) int {
	root := buildDict(words)
	var dfs func(*Trie, int, *int)
	dfs = func(root *Trie, i int, total *int) {
		isLeaf := true
		for _, child := range root.children {
			if child != nil {
				isLeaf = false
				dfs(child, i+1, total)
			}
		}
		if isLeaf {
			*total += i
		}
	}
	ans := 0
	dfs(root, 1, &ans)
	return ans
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"tell", "me", "bell"}))
}
