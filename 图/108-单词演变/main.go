package main

import "fmt"

// 获取word相邻的单词
func getCandidates(word string) []string {
	var res []string
	wordLen := len(word)
	for i := 0; i < 26; i++ {
		for j := 0; j < wordLen; j++ {
			candidate := 'a' + i
			if word[j] != byte(candidate) {
				res = append(res, word[:j]+string(rune(candidate))+word[j+1:])
			}
		}
	}
	return res
}

// 将单词存入hash表
func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, ok := wordMap[word]; !ok {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}
	return wordMap
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap, que, depth := getWordMap(wordList, beginWord), []string{beginWord}, 0
	for len(que) > 0 {
		depth++
		qLen := len(que)
		for i := 0; i < qLen; i++ {
			word := que[0]
			que = que[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, ok := wordMap[candidate]; ok {
					if candidate == endWord {
						return depth + 1
					}
					delete(wordMap, candidate)
					que = append(que, candidate)
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("fuck", "fake", []string{"fack", "fakk", "fake"}))
}
