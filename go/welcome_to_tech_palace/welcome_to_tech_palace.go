package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	upperCustomer := strings.ToUpper(customer)

	return "Welcome to the Tech Palace, " + upperCustomer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starsLine := strings.Repeat("*", numStarsPerLine)
	return starsLine + "\n" + welcomeMsg + "\n" + starsLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msgWithoutStars := strings.ReplaceAll(oldMsg, "*", "")
	trimmedMsg := strings.TrimSpace(msgWithoutStars)

	return trimmedMsg
}
