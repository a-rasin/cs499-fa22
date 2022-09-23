package wordutil

import (
	"strings"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {
	s = strings.ToLower(s)
	m := make(map[string][]int)
	words := strings.Fields(s)

	for _, word := range words {
		cur_str := s[0:]
		cut_length := -1

		if m[word] != nil || len(m[word]) != 0 {
			continue;
		}

		for strings.Contains(cur_str, word) {
			cur_index := strings.Index(cur_str, word)
			cut_length = cut_length + cur_index + 1
			m[word] = append(m[word], cut_length)

			cur_str = cur_str[cur_index + 1:]
		}
	}
	return m
}
