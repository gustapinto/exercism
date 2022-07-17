package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	letters := map[rune]int{}

	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			letters[r] += 1
		}
	}

	return len(letters) == 26
}
