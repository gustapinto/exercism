package anagram

import (
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	for _, candidate := range candidates {
		if CandidateIsEqualToSubject(subject, candidate) ||
			!CandidateHasTheSameNumberOfLetters(subject, candidate) ||
			!CandidateHasOnlyTheSubjectLetters(subject, candidate) {
			continue
		}

		anagrams = append(anagrams, candidate)
	}

	return anagrams
}

func CandidateIsEqualToSubject(subject, candidate string) bool {
	return strings.ToLower(subject) == strings.ToLower(candidate)
}

func CandidateHasTheSameNumberOfLetters(subject, candidate string) bool {
	return len(subject) == len(candidate)
}

func CandidateHasOnlyTheSubjectLetters(subject, candidate string) bool {
	subjectLettersCount := map[rune]int{}
	anagramLettersCount := map[rune]int{}

	for _, r := range subject {
		r = unicode.ToLower(r)
		subjectLettersCount[r] += 1
	}

	for _, r := range candidate {
		r = unicode.ToLower(r)
		anagramLettersCount[r] += 1

		if _, exists := subjectLettersCount[r]; !exists {
			return false
		}
		if anagramLettersCount[r] > subjectLettersCount[r] {
			return false
		}
	}

	return true
}
