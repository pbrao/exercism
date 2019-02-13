package reverse

import (
	"strings"
)

// String takes a string as input and returns its reverse
func String(word string) string {
	var reverse []string
	letters := strings.Split(word, "")
	for x := len(letters) - 1; x >= 0; {
		reverse = append(reverse, letters[x])
		x = x - 1
	}
	return strings.Join(reverse, "")
}
