// Package twofer tests out conditions and string interpolation 
package twofer

import "fmt"

// ShareWith  will return a conditional string based on input
func ShareWith(name string) string {
	var returnName string	
	if name == "" { 
		returnName = "you"
	} else {
		returnName = name
	}
	response := fmt.Sprintf("One for %s, one for me.", returnName)
	return response
}
