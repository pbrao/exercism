package accumulate

// Accumulate applies the input operation to to the input and returns it
func Accumulate(input []string, operation func(string) string) []string {

	var output []string
	var result string
	for _, item := range input {
		result = operation(item)
		output = append(output, result)
	}

	return output
}
