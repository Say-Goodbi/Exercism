package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var border = strings.Repeat("*", numStarsPerLine)
    return border +  "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	starless := strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
	return starless
}
