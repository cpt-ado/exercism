// Package accumulate tranforms a collection with the given operation
package accumulate

// Accumulate tranforms the given collection with the given operation
func Accumulate(input []string, converter func(string) string) []string {
	n := len(input)
	output := make([]string, n)
	for i := 0; i < n; i++ {
		output[i] = converter(input[i])
	}
	return output
}
