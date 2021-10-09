package main

type Trie struct {
	children [26]*Trie
	end      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		if node.children[v-'a'] == nil {
			node.children[v-'a'] = &Trie{}
		}
		node = node.children[v-'a']
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return true
}
