// Package raindrops gets raindrop factors of a number
package raindrops

import "strconv"

// Convert returns a string with corresponding raindrop
// sounds for the number given
func Convert(n int) string {
	result := ""
	if n%3 == 0 {
		result = result + "Pling"
	}
	if n%5 == 0 {
		result = result + "Plang"
	}
	if n%7 == 0 {
		result = result + "Plong"
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}
