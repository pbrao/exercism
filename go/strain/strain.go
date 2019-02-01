package strain

// Ints is a slice of ints
type Ints []int

// Lists is a slice of ints of ints
type Lists [][]int

// Strings is a slice of strings
type Strings []string

// Keep applies a method and returns the items that pass
func (i Ints) Keep(f func(int) bool) Ints {

	var result Ints

	for _, item := range i {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Discard applies a method and returns the items that do NOT pass
func (i Ints) Discard(f func(int) bool) Ints {

	var result Ints
	for _, item := range i {
		if !f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Keep applies a method and returns the items that pass
func (l Lists) Keep(f func([]int) bool) Lists {

	var result Lists
	for _, item := range l {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Keep applies a method and returns the items that pass
func (s Strings) Keep(f func(string) bool) Strings {

	var result Strings
	for _, item := range s {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
