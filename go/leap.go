package leap

func IsLeapYear(year int) bool {
	isDivisibleByFour := year%4 == 0
	isDivisibleByOneHundred := year%100 == 0
	isDivisibleByFourHundred := year%400 == 0

	return isDivisibleByFour && !isDivisibleByOneHundred || isDivisibleByFourHundred
}
