package luhn

import (
	"strings"
	"unicode"
)

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Valid(id string) bool {
	sum := 0
	reversedId := strings.ReplaceAll(reverse(id), " ", "")

	if len(reversedId) <= 1 {
		return false
	}

	for i, r := range reversedId {
		if !unicode.IsDigit(r) {
			return false
		}

		value := int(r - '0') // Converte rune para int

		if i%2 == 1 { // Dobra o segundo valor da esquerda para a direita, pois invertemos o id
			value = value * 2

			if value > 9 {
				value -= 9
			}
		}

		sum += value
	}

	return sum%10 == 0
}
