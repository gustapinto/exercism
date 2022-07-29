package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	score := map[string]int{}

	for value, letters := range in {
		for _, letter := range letters {
			lowercaseLetter := strings.ToLower(letter)

			score[lowercaseLetter] += value
		}
	}

	return score
}
