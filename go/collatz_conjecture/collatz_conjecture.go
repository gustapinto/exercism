package collatzconjecture

import "errors"

func IsEven(n int) bool {
	return n%2 == 0
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("should use a positive integer")
	}

	count := 0

	for n != 1 {
		count += 1

		if IsEven(n) {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}

	return count, nil
}
