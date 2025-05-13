package zdutil

import "unicode"

// IsLetter checks if all characters in the string are letters.
// It returns true if every character in the string is a letter,
// otherwise it returns false.
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsNumber checks if all characters in the string are numbers.
// It returns true if every character in the string is a number,
// otherwise it returns false.
func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}
