package strand

import "strings"

// ToRNA converts a DNA strand to an RNA strand
func ToRNA(dna string) string {

	var sb strings.Builder

	for _, letter := range dna {

		/*
					switch letter {
					case 'G':
						rna += "C"
					case 'C':
						rna += "G"
					case 'T':
						rna += "A"
					case 'A':
						rna += "U"
					}
			$ go test -v --bench . --benchmem
			=== RUN   TestRNATranscription
			--- PASS: TestRNATranscription (0.00s)
			    rna_transcription_test.go:11: PASS: Empty RNA sequence
			    rna_transcription_test.go:11: PASS: RNA complement of cytosine is guanine
			    rna_transcription_test.go:11: PASS: RNA complement of guanine is cytosine
			    rna_transcription_test.go:11: PASS: RNA complement of thymine is adenine
			    rna_transcription_test.go:11: PASS: RNA complement of adenine is uracil
			    rna_transcription_test.go:11: PASS: RNA complement
			goos: linux
			goarch: amd64
			pkg: github.com/pbrao/exercism/go/rna-transcription
			BenchmarkRNATranscription-4      3000000               515 ns/op              96 B/op            11 allocs/op
			PASS
			ok      github.com/pbrao/exercism/go/rna-transcription  2.066s
		*/

		switch letter {
		case 'G':
			sb.WriteString("C")
		case 'C':
			sb.WriteString("G")
		case 'T':
			sb.WriteString("A")
		case 'A':
			sb.WriteString("U")
		}
		/*
					$ go test -v --bench . --benchmem
			=== RUN   TestRNATranscription
			--- PASS: TestRNATranscription (0.00s)
			    rna_transcription_test.go:11: PASS: Empty RNA sequence
			    rna_transcription_test.go:11: PASS: RNA complement of cytosine is guanine
			    rna_transcription_test.go:11: PASS: RNA complement of guanine is cytosine
			    rna_transcription_test.go:11: PASS: RNA complement of thymine is adenine
			    rna_transcription_test.go:11: PASS: RNA complement of adenine is uracil
			    rna_transcription_test.go:11: PASS: RNA complement
			goos: linux
			goarch: amd64
			pkg: github.com/pbrao/exercism/go/rna-transcription
			BenchmarkRNATranscription-4     10000000               216 ns/op              56 B/op            6 allocs/op
			PASS
			ok      github.com/pbrao/exercism/go/rna-transcription  2.407s
		*/
	}

	return sb.String()
}
