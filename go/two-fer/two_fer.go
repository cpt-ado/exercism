// Package twofer produces a string with the
// name given by the user
package twofer

// ShareWith returns the string
// "One for X, one for me." given
// a string "X"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
