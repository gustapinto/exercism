package prime

func Factors(n int64) []int64 {
	factors := []int64{}

	prime := int64(2) // Obtém o primeiro número primo, vulgo 2
	for n > 1 && n >= prime {
		if n%prime == 0 {
			// Se N for divisivel pelo número primo adiciona o fator e divide por N
			factors = append(factors, prime)
			n /= prime
		} else {
			// Senão for divisivel pelo fator (e consequentemente pelos seus produtos) incrementa 1
			// e tenta novamente
			prime += 1
		}
	}

	return factors
}
