package utils

// IsAlphanumericRune checks if a given Unicode rune is alphanumeric
func IsAlphanumericRune(c rune) bool {
	return ('0' <= c && c <= '9') || ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

// CapitalizeUnicode returns the capitalized unicode character of a lowercase letter. Otherwise return the same character
func CapitalizeUnicode(c rune) rune {
	if !('a' <= c && c <= 'z') {
		return c
	}

	// 32 is ASCII difference between lowercase letter and an uppercase letter. Return the uppercase letter of this lowercase letter
	return c - 32
}
