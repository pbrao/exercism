package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

// Frequency ...
type Frequency map[string]int

// WordCount ...
func WordCount(input string) Frequency {
	result := make(Frequency)
	fmt.Println("INPUT:  ", input)
	// Remove all non-alphanumeric characters
	//reg, _ := regexp.Compile("[^a-zA-Z0-9,]+")
	//processedString := reg.ReplaceAllString(input, "")
	re := regexp.MustCompile("\n|[^a-zA-Z0-9\\s,']+")
	processedString := strings.ToLower(re.ReplaceAllString(input, ""))
	fmt.Println("PROCESSEDSTRING:  ", processedString)
	words := regexp.MustCompile("[,\\s]+").Split(processedString, -1)
	fmt.Println("WORDS:  ", words)

	for _, word := range words {
		fmt.Println("WORD:  ", word)
		result[word] = result[word] + 1
	}
	fmt.Println("RESULT:  ", result)
	return result
}
