package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strArray := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	wordsCounter(strArray)
}

func wordsCounter(words []string) map[string]int {
	wordsMapWithA := make(map[string]int)
	wordsMap := make(map[string]int)

	for _, word := range words {
		var count int
		for _, letter := range word {
			if strings.Contains(string(letter), "a") {
				count = count + 1
			}
		}
		if count > 0 {
			wordsMapWithA[word] = count
		} else {
			wordsMap[word] = len(word)
		}

	}

	sortWords(wordsMapWithA)
	sortWords(wordsMap)

	return wordsMapWithA

}
func sortWords(words map[string]int) {
	type kv struct {
		Key   string
		Value int
	}

	var sortSlice []kv
	for k, v := range words {
		sortSlice = append(sortSlice, kv{k, v})
	}

	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].Value > sortSlice[j].Value
	})

	for _, kv := range sortSlice {
		fmt.Printf("%s,", kv.Key)
	}
}
