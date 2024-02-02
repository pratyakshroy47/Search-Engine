package utils

import (
	"strings"
	"unicode"
)

// tokenize breaks down the input text into a slice of tokens.
func tokenize(text string) []string {
	// FieldsFunc splits the text into words based on the given predicate function.
	return strings.FieldsFunc(text, func(r rune) bool {
		// Tokens are comprised of letters and numbers; anything else is considered a separator.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// analyze processes the input text by tokenizing, converting to lowercase, removing stopwords, and applying stemming.
func analyze(text string) []string {
	// Tokenize the input text into individual words.
	tokens := tokenize(text)
	// Apply filters sequentially to process the tokens.
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}