package worldhelper

import (
	"regexp"
	"strings"
)

// WordCensor replaces profane words with asterisks in a given text.
func WordCensor(input string) string {
	// Define a list of profane words to censor
	profaneWords := []string{"prayuth", "fuck", "bitch", "Piss off",
		"Fuck ",
		"FucK ",
		"fUck ",
		"fUCk ",
		"Prayuth",
		"PraYuth",
		"prayutH",
		"PrAyuth",
		"Piss Off",
		"Piss 0ff",
		"PIss 0ff",
		"P1ss Off",
		"Plss Off"}

	// Replace profane words with asterisks
	for _, word := range profaneWords {
		// Construct a regular expression to match the word surrounded by word boundaries
		re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(word) + `\b`)
		input = re.ReplaceAllString(input, strings.Repeat("*", len(word)))
	}

	return input
}

// BuffaloNeckGenerator generates a string consisting of 'ค' repeated n times.
// If n is negative or zero, it returns a string with one 'ค'.
func BuffaloNeckGenerator(n int) string {
	if n <= 0 {
		return "ค"
	}

	return strings.Repeat("ค", n)
}
