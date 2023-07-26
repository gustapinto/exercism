package reverse

func Reverse(input string) string {
	// Usa uma slice de ruens para lidar com caracteres que não são UTF-8
	reversed := []rune(input)

	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return string(reversed)
}
