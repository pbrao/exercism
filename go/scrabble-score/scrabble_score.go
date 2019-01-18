package scrabble

import (
	"strings"
)

// Score calculates the numeric Scrabble score and returns it
func Score(word string) int {

	score := 0

	word = strings.ToLower(word)

	for i, letter := range word {
		switch {
		case strings.ContainsAny(string(letter), "a || e || i || o || u || l || n || r || s || t"):
			score++
		case strings.ContainsAny(string(letter), "d || g"):
			score += 2
		case strings.ContainsAny(string(letter), "b || c || m || p"):
			score += 3
		case strings.ContainsAny(string(letter), "f || h || v || w || y"):
			score += 4
		case strings.ContainsAny(string(letter), "k"):
			score += 5
		case strings.ContainsAny(string(letter), "j || x"):
			score += 8
		case strings.ContainsAny(string(letter), "q || z"):
			score += 10
		}

		_ = i
	}

	return score
}
