package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children [26]*Trie
	end      bool
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

func (t *Trie) SearchPrefix(prefix string) string {
	node := t
	s := strings.Builder{}
	for _, v := range prefix {
		if node.end || node.children[v-'a'] == nil {
			break
		}
		s.WriteRune(v)
		node = node.children[v-'a']
	}
	if node.end {
		return s.String()
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	t := Trie{}
	for _, v := range dictionary {
		t.Insert(v)
	}
	words := strings.Split(sentence, " ")
	for i, v := range words {
		prefix := t.SearchPrefix(v)
		if prefix != "" {
			words[i] = prefix
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "rat", "bat"}, "the cattle was rattled by the battery"))
}
