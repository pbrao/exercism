package proverb

import "fmt"

// Proverb constructs a rhyme based on input elements
func Proverb(rhyme []string) []string {

	var result []string
	// If no data is passed in, return an empty resultset.
	if len(rhyme) == 0 {
		return []string{}
	}

	// If 2 or more elements are passed in, then construct the rhyme.
	if len(rhyme) > 1 {
		var i int
		for i = 0; i < len(rhyme)-1; i++ {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	// Append the last line of the rhyme to the result
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return result
}
