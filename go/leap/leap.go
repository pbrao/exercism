// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear calculates whether a year is a leap year or no
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
