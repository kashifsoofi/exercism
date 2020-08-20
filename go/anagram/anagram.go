// Package anagram implements a simple library that
// determine the anagrams of a word in a given of candidates
package anagram

import (
	"strings"
	"unicode"
)

// Detect given word and candidates, detects and returns the anagrams of the word
func Detect(word string, candidates []string) []string {
	c := map[rune]int{}
	for _, v := range word {
		k := unicode.ToLower(v)
		c[k]++
	}

	anagrams := []string{}
	for _, s := range candidates {
		if areAnagram(word, s, c) {
			anagrams = append(anagrams, s)
		}
	}

	return anagrams
}

func areAnagram(s1, s2 string, c map[rune]int) bool {
	if len(s1) != len(s2) {
		return false
	}

	if strings.EqualFold(s1, s2) {
		return false
	}

	c1 := map[rune]int{}
	for _, v := range s2 {
		k := unicode.ToLower(v)
		_, ok := c[k]
		if !ok {
			return false
		}

		c1[k]++
		if c1[k] > c[k] {
			return false
		}
	}

	for k := range c {
		if c[k] != c1[k] {
			return false
		}
	}

	return true
}
