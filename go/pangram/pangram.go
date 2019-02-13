package pangram

import (
	"unicode"
)

// IsPangram determines whether the string is a pangram
func IsPangram(sentence string) bool {

	unique := make(map[rune]rune)
	for _, letter := range sentence {
		if unicode.IsLetter(letter) {
			lowerLetter := unicode.ToLower(letter)
			unique[lowerLetter] = letter
		}
	}

	if len(unique) != 26 {
		return false
	}
	return true
}
