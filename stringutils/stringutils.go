package stringutils

// Reverse returns the reversed string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsEmpty returns true if the string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}
