// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)
	words := strings.Fields(s)
	var acronym string
	for _, word := range words {
		acronym = acronym + string(word[0])
	}
	return strings.ToUpper(acronym)
}
