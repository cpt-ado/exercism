// Package twofer produces a string with the
// name given by the user
package twofer

// ShareWith returns the string
// "One for X, one for me." given
// a string "X"
func ShareWith(name string) string {
	if name != "" {
	return "One for " + name + ", one for me."
	} else {
	return "One for you, one for me."
	}
}
