// Package twofer tests out conditions and string interpolation
package twofer

// ShareWith  will return a conditional string based on input
func ShareWith(name string) string {
	var returnName string
	var response string
	if name == "" {
		returnName = "you"
	} else {
		returnName = name
	}
	//  response := fmt.Sprintf("One for %s, one for me.", returnName)
	response = "One for " + returnName + ", one for me."
	return response
}
