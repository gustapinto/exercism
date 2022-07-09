package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	decimal, err := ToDecimal(inputBase, inputDigits)
	if err != nil {
		return []int{}, err
	}

	digits := ToBase(outputBase, decimal)
	if len(digits) == 0 {
		return []int{0}, nil
	}

	return digits, nil
}

func ToDecimal(b int, digits []int) (int, error) {
	n := len(digits)
	decimal := 0
	for _, digit := range digits {
		if 0 <= digit && digit < b {
			n -= 1
			decimal += digit * int(math.Pow(float64(b), float64(n)))
		} else {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	return decimal, nil
}

func ToBase(b, n int) []int {
	m := []int{}

	for n != 0 {
		m = append([]int{n % b}, m...)
		n /= b
	}

	return m
}
