package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "Hello chonnal"
	var bad_word = []string{"Hell", "chon"}
	var result = word_censor(text, bad_word)
	fmt.Println(result)

	result = word_censor("Hello Hello Kitty", []string{"lo"})
	fmt.Println(result)
}

func word_censor(text string, bad_words []string) string {
	for _, word := range bad_words {
		text = strings.ReplaceAll(text, word, strings.Repeat("*", len(word)))
	}
	// for i := 0; i < len(bad_words); i++ {
	//     var word = bad_words[i]
	//     text = strings.ReplaceAll(text, word, strings.Repeat("*", len(word)))
	// }
	return text
}
