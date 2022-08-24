package wordcount

import (
	"strings"
	"regexp"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}

	// Usa regex para encontrar todas as palavras
	re := regexp.MustCompile(`[a-z0-9]+(['][a-z0-9]+)?`)
	words := re.FindAllString(strings.ToLower(phrase), -1)

	for _, word := range words {
		freq[word]++
	}

	return freq
}
