package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	if word == "" || len(word) == 1 {
		return true
	}

	word = strings.ToLower(word)
	letters := make(map[rune]int, len(word))

	for _, l := range word {
		if unicode.IsPunct(l) || unicode.IsSpace(l) {
			continue
		}

		if letters[l] != 0 {
			return false
		}

		letters[l] += 1
	}

	return true
}
