package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = make(Histogram)
	var A, C, G, T int

	for _, letter := range d {
		switch letter {
		case 'A':
			A++
		case 'C':
			C++
		case 'G':
			G++
		case 'T':
			T++
		default:
			return h, errors.New("Invalid value")
		}
	}
	h['A'] = A
	h['C'] = C
	h['G'] = G
	h['T'] = T

	return h, nil
}
