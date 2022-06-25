package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	// Usa regex para splittar a string usando multiplos delimitadores
	words := regexp.MustCompile("[\\-\\ \\_\\s]+").Split(s, -1)

	acronym := ""
	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}

	return acronym
}
