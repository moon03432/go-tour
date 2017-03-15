package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

/*
 * Exercice on Map
 *
 * Implement WordCount. It should return a map of the counts of each “word” in the string s.
 * The wc.Test function runs a test suite against the provided function and prints success or failure.
 */
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _,w := range strings.Fields(s) {
		_,ok := m[w]
		if ok {
			m[w] += 1
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}