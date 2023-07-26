package sieve

func Sieve(limit int) []int {
	// Inicializa uma slice de booleanos que vão marcar quais números são primos e quais não são
	prime := make([]bool, limit+1)
	for i := range prime {
		prime[i] = true
	}

	// Percorre a slice de booleanos marcando os números primos
	for i := 2; i*i < limit; i++ {
		if prime[i] == true { // Percorre todos os números marcados como não primos
			for j := i * 2; j <= limit; j += i {
				prime[j] = false // Marca os números não primos
			}
		}
	}

	// Com base nos números marcados como primos gera uma slice de inteiros
	primeNumbers := make([]int, 0)
	for number, isPrime := range prime {
		if isPrime && number >= 2 {
			primeNumbers = append(primeNumbers, number)
		}
	}

	return primeNumbers
}
