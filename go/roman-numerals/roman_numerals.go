package romannumerals

import "errors"

// ToRomanNumeral converts decimal numbers into roman numerals
func ToRomanNumeral(arabic int) (string, error) {

	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("true")
	}

	var romannumeral string
	residual := arabic

	for a := 0; a < (arabic / 1000); a++ {
		romannumeral = romannumeral + "M"
		residual = residual - 1000
	}
	arabic = residual

	if arabic >= 900 {
		romannumeral = romannumeral + "CM"
		residual = residual - 900
	} else {
		for a := 0; a < (arabic / 500); a++ {
			romannumeral = romannumeral + "D"
			residual = residual - 500
		}
	}
	arabic = residual

	if arabic >= 400 {
		romannumeral = romannumeral + "CD"
		residual = residual - 400
	} else {
		for a := 0; a < (arabic / 100); a++ {
			romannumeral = romannumeral + "C"
			residual = residual - 100
		}
	}
	arabic = residual

	if arabic >= 90 {
		romannumeral = romannumeral + "XC"
		residual = residual - 90
	} else {
		for a := 0; a < (arabic / 50); a++ {
			romannumeral = romannumeral + "L"
			residual = residual - 50
		}
	}
	arabic = residual

	if arabic >= 40 {
		romannumeral = romannumeral + "XL"
		residual = residual - 40
	} else {
		for a := 0; a < (arabic / 10); a++ {
			romannumeral = romannumeral + "X"
			residual = residual - 10
		}
	}
	arabic = residual

	if arabic == 9 {
		romannumeral = romannumeral + "IX"
		residual = residual - 9
	} else {
		for a := 0; a < (arabic / 5); a++ {
			romannumeral = romannumeral + "V"
			residual = residual - 5
		}
	}
	arabic = residual

	if arabic == 4 {
		romannumeral = romannumeral + "IV"
	} else {
		for a := 0; a < (arabic / 1); a++ {
			romannumeral = romannumeral + "I"
			residual = residual - 1
		}
	}
	return romannumeral, nil
}
