package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid number")
	}

	total := uint64(1)

	for i := 1; i < number; i++ {
		total *= 2
	}

	return total, nil
}

func Total() uint64 {
	total := uint64(0)

	for i := 1; i <= 64; i++ {
		squareTotal, _ := Square(i)
		total += squareTotal
	}

	return total
}
