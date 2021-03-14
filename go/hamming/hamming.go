// Package hamming calculates the hamming distance
// between two strings
package hamming

import "errors"

// Distance returns the hamming distance between the 
// strings a and b, or an error if lengths do not match
func Distance(a, b string) (int, error) {
	n := len(a)
	if n != len(b) {
		return 0, errors.New("length of strings does not match")
	}
        distance := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
