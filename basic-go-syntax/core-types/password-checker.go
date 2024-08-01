package coretypes

import (
	"fmt"
	"unicode"
)

func PasswordChecker(s string) {
	// This is a password checker function.
	// This function initially checks the length of the password. The minimnum is 8
	// Then, the function checks each string character converted into rune type
	// For every rune type, if the all uppercase, lowercase, number, and symbol exist in the password, the password is valid
	// If none exists, the password is invalid

	var hasUpper bool
	var hasLower bool
	var hasSymbol bool
	var hasNumber bool
	var hasValidLen bool

	if len(s) >= 8 {
		hasValidLen = true
	}

	// The loop must range over string then converted into rune character by character
	// Otherwise, error happens if the string is converted into rune type first
	for _, c := range s {

		if unicode.IsUpper(c) {
			hasUpper = true
		}

		if unicode.IsLower(c) {
			hasLower = true
		}

		if unicode.IsNumber(c) {
			hasNumber = true
		}

		if unicode.IsSymbol(c) || unicode.IsPunct(c) {
			hasSymbol = true
		}
	}

	if hasUpper && hasLower && hasNumber && hasSymbol && hasValidLen {
		fmt.Println("Your password is good!")

	} else {
		fmt.Println("Your password does not comply with the requirements, please try again.")
	}
}
