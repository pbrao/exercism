package hamming

import "errors"

// Distance calculates the Hamming distance
// between 2 DNA strands
// If they aren't the same length, it returns an error
func Distance(a, b string) (int, error) {

	// Check if DNA strings are the same length
	if len(a) != len(b) {
		return 0, errors.New("true")
	}

	matchCount := 0
	for i := 0; i < len(a); i++ {

		if a[i] != b[i] {
			matchCount++
		}
	}

	return matchCount, nil
}
