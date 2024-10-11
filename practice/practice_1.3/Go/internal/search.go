package search

import (
	"strings"
)

// Get reversed string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Prefix computation function
func computePrefixFunction(pattern string) []int {
	prefix := make([]int, len(pattern))
	k := 0

	for i := 1; i < len(pattern); i++ {
		for k > 0 && pattern[k] != pattern[i] {
			k = prefix[k-1]
		}
		if pattern[k] == pattern[i] {
			k++
		}
		prefix[i] = k
	}
	return prefix
}

// Knuth-Morris-Pratt algorithm func
func KMPSearch(text string, pattern string, caseSensitivity bool, method string, count int) []int {
	// Если caseSensitivity флаг выключен, привести строки к нижнему регистру
	if !caseSensitivity {
		text = strings.ToLower(text)
		pattern = strings.ToLower(pattern)
	}

	if method == "last" {
		text = reverseString(text)
		pattern = reverseString(pattern)
	}

	// Maximum number of founded sub strings is count
	counter := 1

	// Применение алгоритма КМП
	prefix := computePrefixFunction(pattern)
	matches := []int{}
	n := len(text)
	m := len(pattern)
	k := 0

	for i := 0; i < n; i++ {
		for k > 0 && pattern[k] != text[i] {
			k = prefix[k-1]
		}
		if pattern[k] == text[i] {
			k++
		}
		if k == m {
			if counter > count {
				break
			}
			if method == "last" {
				matches = append(matches, n-i-1)
			} else {
				matches = append(matches, i-m+1)	
			}
			k = prefix[k-1]
			counter++
		}
	}

	return matches
}

// Обёртка для поиска подстрок
func Search(text string, subStrings []string, caseSensitivity bool, method string, count int) map[string][]int {
	results := make(map[string][]int)

	for _, subString := range subStrings {
		results[subString] = KMPSearch(text, subString, caseSensitivity, method, count)
	}

	return results
}
