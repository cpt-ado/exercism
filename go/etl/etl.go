// Package etl updates old scrabble data into new format
package etl

import "strings"

// Transform inverts the shape of the given map from value to slice
// of letters into a map from letters to values
func Transform (input map[int][]string) map[string]int {
	output := make(map[string]int)
	for v, letters := range input {
		for l := 0; l < len(letters); l++ {
			output[strings.ToLower(string(letters[l]))] = v
		}
	}
	return output
}
