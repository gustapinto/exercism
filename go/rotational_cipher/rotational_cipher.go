package rotationalcipher

import (
	"unicode"
)

const (
	minLowerRuneVal = int(rune('a'))
	maxLowerRuneVal = int(rune('z'))
	minUpperRuneVal = int(rune('A'))
	maxUpperRuneVal = int(rune('Z'))
)

func ShiftRune(val, min, max int) rune {
	if val > max {
		return rune(min + (val - max) - 1)
	}

	return rune(val)
}

func RotationalCipher(plain string, shiftKey int) (ciphered string) {
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	for _, d := range plain {
		if !unicode.IsLetter(d) {
			ciphered += string(d)
			continue
		}

		val := int(d) + shiftKey

		if unicode.IsLower(d) {
			ciphered += string(ShiftRune(val, minLowerRuneVal, maxLowerRuneVal))
			continue
		}

		ciphered += string(ShiftRune(val, minUpperRuneVal, maxUpperRuneVal))
	}

	return ciphered
}
