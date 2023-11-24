package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(word1, word2 string) bool {
	word1Chars := strings.Split(word1, "")
	word2Chars := strings.Split(word2, "")
	sort.Strings(word1Chars)
	sort.Strings(word2Chars)

	sortedWord1 := strings.Join(word1Chars, "")
	sortedWord2 := strings.Join(word2Chars, "")

	return sortedWord1 == sortedWord2
}

func validateInput(word string) bool {
	if len(word) > 10 {
		return false
	}

	for _, char := range word {
		if char == ' ' || strings.ContainsAny(string(char), "!@#$%^&*()-+[]{};:'\"<>?/.,") {
			return false
		}
	}

	return true
}

func checkAnagram(word1, word2 string) {
	if !validateInput(word1) || !validateInput(word2) {
		panic("Input tidak valid")
	}

	if isAnagram(word1, word2) {
		fmt.Printf("[%s] & [%s] merupakan anagram\n", word1, word2)
	} else {
		fmt.Printf("[%s] & [%s] bukan merupakan anagram\n", word1, word2)
	}
}
