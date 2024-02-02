package utils

import (
	"strings"
	snowballeng "github.com/kljensen/snowball/english"
)

// lowercaseFilter converts all tokens to lowercase.
func lowercaseFilter(tokens []string) []string {
	// Create a new slice to store lowercase tokens.
	result := make([]string, len(tokens))
	for i, token := range tokens {
		result[i] = strings.ToLower(token)
	}
	return result
}

// stopwordFilter removes common stopwords from the list of tokens.
func stopwordFilter(tokens []string) []string {
	// Define a set of stopwords for efficient lookup.
	stopwords := map[string]struct{}{
		"a":    {}, "and": {}, "have": {}, "i": {}, "in": {},
		"of":   {}, "that": {}, "the": {}, "to": {},
	}
	// Create a new slice to store non-stopword tokens.
	result := make([]string, 0, len(tokens))
	// Iterate through tokens, adding non-stopwords to the result slice.
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			result = append(result, token)
		}
	}
	return result
}

// stemmerFilter applies stemming to reduce tokens to their base or root form.
func stemmerFilter(tokens []string) []string {
	// Create a new slice to store stemmed tokens.
	result := make([]string, len(tokens))
	// Iterate through tokens, applying the snowball English stemming algorithm.
	for i, token := range tokens {
		result[i] = snowballeng.Stem(token, false)
	}
	return result
}
