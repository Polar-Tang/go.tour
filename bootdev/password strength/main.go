package main

import (
	"regexp"
)

func isValidPassword(password string) bool {
	digitRegex := regexp.MustCompile(`\d`)        // Matches any digit
	uppercaseRegex := regexp.MustCompile(`[A-Z]`) // Matches any uppercase letter
	checkLength := len(password) >= 5 && len(password) <= 12
	hasDigit := digitRegex.MatchString(password)
	hasUppercase := uppercaseRegex.MatchString(password)

	return hasDigit && hasUppercase && checkLength

}
