package main

import "fmt"

type Trie struct {
	children [26]*Trie
	end      bool
}

type MagicDictionary struct {
	root *Trie
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		root: &Trie{},
	}
}

func (dict *MagicDictionary) BuildDict(dictionary []string) {
	var node *Trie
	for _, v := range dictionary {
		node = dict.root
		for _, ch := range v {
			if node.children[ch-'a'] == nil {
				node.children[ch-'a'] = &Trie{}
			}
			node = node.children[ch-'a']
		}
		node.end = true
	}
}

// edit表示修改的字符数量
func searchCore(root *Trie, searchWord string, i, edit int) bool {
	if root == nil {
		return false
	}
	if root.end && i == len(searchWord) && edit == 1 {
		return true
	}
	var (
		found bool
		next  int //循环中edit不应该赋值，不然后续全会被视为已修改字符
	)
	if i < len(searchWord) && edit <= 1 {
		found = false
		for j := 0; j < 26 && !found; j++ {
			if j != int(searchWord[i]-'a') {
				next = edit + 1
			} else {
				next = edit
			}
			found = searchCore(root.children[j], searchWord, i+1, next)
		}
		return found
	}
	return false
}

func (dict *MagicDictionary) Search(searchWord string) bool {
	return searchCore(dict.root, searchWord, 0, 0)
}

func main() {
	t := Constructor()
	t.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(t.Search("hhllo"))
}
