package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be greater than 1")
	}

	limit := n
	primes := []int{}
	for len(primes) < n {
		limit *= 2
		primes = Sieve(limit)
	}

	return primes[n-1], nil
}

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
