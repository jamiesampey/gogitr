package palidrome

import (
	"unicode"
)

func IsPalidrome(input string) bool {
	var runes []rune

	for _, c := range input {
		if unicode.IsLetter(c) {
			runes = append(runes, unicode.ToLower(c))
		}
	}

	var lettersLen = len(runes)
	for i := 0; i < lettersLen/2; i++ {
		if runes[i] != runes[lettersLen-i-1] {
			return false
		}
	}

	return true
}
