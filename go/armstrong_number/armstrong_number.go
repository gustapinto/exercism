package armstrong

import (
	"fmt"
	"strconv"
)

func pow(x, y int) int {
	if x == 0 {
		return 1
	}
	if y == 1 {
		return x
	}

	n := x
	for i := 2; i <= y; i++ {
		n *= x
	}
	return n
}

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	digits := len(s)
	if digits == 1 {
		return true
	}

	sum := 0
	for _, r := range s {
		// Usa uma característica especial de runes para converter o caractere (r) em int, subtraindo
		// a rune '0' temos acesso ao valor UTF-8 que o caractere (r) representa, podendo portanto
		// converter para inteito
		number := int(r - '0')
		sum += pow(number, digits)
		fmt.Print(number, digits, " ")
	}

	return n == sum
}
