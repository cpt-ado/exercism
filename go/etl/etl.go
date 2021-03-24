package etl

import "strings"

func Transform (input map[int][]string) map[string]int {
	output := make(map[string]int)
	for v, letters := range input {
		for l := 0; l < len(letters); l++ {
			output[strings.ToLower(string(letters[l]))] = v
		}
	}
	return output
}
