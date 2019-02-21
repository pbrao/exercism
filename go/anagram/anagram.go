package anagram

import (
	"bytes"
	"sort"
	"strings"
)

// Detect determines whether there are any anagrams
func Detect(subject string, candidates []string) []string {

	var matches = make([]string, 0)
	lowerSubject := strings.ToLower(subject)
	letters := strings.Split(lowerSubject, "")
	sort.Strings(letters)
	letterslen := len(letters)
	lettersString := strings.Join(letters, "")
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if letterslen == len(candidate) && lowerSubject != lowerCandidate {

			candidateletters := strings.Split(lowerCandidate, "")
			sort.Strings(candidateletters)
			candidateString := strings.Join(candidateletters, "")
			if bytes.Equal([]byte(lettersString), []byte(candidateString)) {
				matches = append(matches, candidate)
			}
		}
	}
	return matches
}
