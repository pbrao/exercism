package isogram

import (
	"unicode"
)

// IsIsogram determines if the word is an isogram (i.e. has duplicate characters) or not
func IsIsogram(word string) bool {

	var seen = map[rune]bool{}
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}
		lowerLetter := unicode.ToLower(letter)

		if seen[lowerLetter] {
			return false
		}
		seen[lowerLetter] = true
	}
	return true
}
