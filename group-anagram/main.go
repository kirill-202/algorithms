package main

import (
	"fmt"
	"sort"
)

/*
Write a Go function that takes an array of strings and groups the anagrams together. 
An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
typically using all the original letters exactly once.

*/

type Word struct {
	baseWord string
	sortedWord string
}

func (w *Word) sortWord() {
	runes := []rune(w.baseWord)
	fmt.Println("Unsorted runes", runes)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	fmt.Println("sorted runes", runes)
	w.sortedWord = string(runes)
}

func groupAnagrams(anagrams []string) (sorted [][]string) {
	occurences := make(map[string][]string)
	for _, sequence := range anagrams {
		word := Word{
			baseWord: sequence,
		}
		word.sortWord()

		occurences[word.sortedWord] = append(occurences[word.sortedWord], word.baseWord)
	}

	for _, words := range occurences {
		sorted = append(sorted, words)
	}
	return

}



func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	output := groupAnagrams(input) 
	fmt.Printf("The output of the funnction %v", output)
}
// Expected output: [["eat","tea","ate"],["tan","nat"],["bat"]]
