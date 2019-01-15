package raindrops

import "strconv"

// Convert evaluates an integer and
// responds indicating whether it is divisible
// by 3,5, 7
func Convert(number int) string {

	var response string

	if number%3 == 0 {
		response = "Pling"
	}
	if number%5 == 0 {
		response += "Plang"
	}
	if number%7 == 0 {
		response += "Plong"
	}

	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		response = strconv.Itoa(number)
	}

	return response
}
