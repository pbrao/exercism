package scrabble

import "unicode"

// Score calculates the numeric Scrabble score and returns it
func Score(word string) int {

	score := 0

	for i, letter := range word {
		letter = unicode.ToLower(letter)
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}

		_ = i
	}

	return score
}
