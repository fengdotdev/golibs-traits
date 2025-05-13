package handyfuncs

import "strings"

// LookupStringIn checks if a string is present in another string
// It returns true if the term is found in the whole string, otherwise false
// It is case insensitive
func LookupStringIn(term string, whole string) bool {
	if term == "" || whole == "" {
		return true
	}

	result := strings.Contains(strings.ToLower(whole), strings.ToLower(term))
	return result
}
