package main

type Trie struct {
	children [26]*Trie
	val      int
}

type MapSum struct {
	root *Trie
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		root: &Trie{},
	}
}

func (m *MapSum) Insert(key string, val int) {
	node := m.root
	for _, v := range key {
		if node.children[v-'a'] == nil {
			node.children[v-'a'] = &Trie{}
		}
		node = node.children[v-'a']
	}
	node.val = val
}

func sum(node *Trie) int {
	if node == nil {
		return 0
	}
	ans := node.val
	for _, child := range node.children {
		if child != nil {
			ans += sum(child)
		}
	}
	return ans
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for _, v := range prefix {
		if node.children[v-'a'] == nil {
			return 0
		}
		node = node.children[v-'a']
	}
	return sum(node)
}
