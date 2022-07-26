package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	validPrefixes := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}

	for _, validPrefix := range validPrefixes {
		// ^ -> Indica que a expressão buscada desve estar no começo da linha
		re := regexp.MustCompile(`^\[` + validPrefix + `\]`)

		if re.MatchString(text) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	// (...)  -> Monta um grupo de caracteres
	// |      -> Operador "ou"
	// *      -> Operador "um ou muitos"
	// (a|b)* -> busca por uma ou várias ocorrências de "a" ou "b"
	re := regexp.MustCompile(`<(\*|~|=|-)*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (counter int) {
	// .*   -> Operador "qualquer caracter válido"
	// (?i) -> Operador que identifica uma palavra inteira
	re := regexp.MustCompile(`".*(?i)password.*"`)

	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`(?i)end-of-line[0-9]*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (newLines []string) {
	re := regexp.MustCompile(`User\s+(\w+)`)

	for _, line := range lines {
		user := re.FindStringSubmatch(line)
		if len(user) > 1 {
			line = "[USR] " + user[1] + " " + line
		}

		newLines = append(newLines, line)
	}

	return newLines
}
