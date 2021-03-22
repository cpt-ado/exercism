// Package dna defines the DNA type and its methods
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its frequency
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA
// Returns an error if d contains an invalid nucleotide
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	for i := 0; i < len(d); i++ {
		nucleotide := rune(d[i])
		if strings.ContainsRune("ACGT", nucleotide) {
			h[nucleotide]++
		} else {
			return nil, errors.New("dna contains invalid nucleotides")
		}
	}
	return h, nil
}
