package isogram

import (
	"unicode"
)

// IsIsogram determines if the word is an isogram (i.e. has duplicate characters) or not
func IsIsogram(word string) bool {

	counter := make(map[rune]int)
	if word == "" {
		return true
	}
	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}
		lowerLetter := unicode.ToLower(letter)
		counter[lowerLetter] = counter[lowerLetter] + 1
		if counter[lowerLetter] > 1 {
			return false
		}
	}
	return true
}
