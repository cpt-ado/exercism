// Package scrabble calculates the score of words
package scrabble

import "strings"

var value = map[string]int{
	// value 1
	"a": 1,
	"e": 1,
	"i": 1,
	"o": 1,
	"u": 1,
	"l": 1,
	"n": 1,
	"r": 1,
	"s": 1,
	"t": 1,
	// value 2
	"d": 2,
	"g": 2,
	// value 3
	"b": 3,
	"c": 3,
	"m": 3,
	"p": 3,
	// value 4
	"f": 4,
	"h": 4,
	"v": 4,
	"w": 4,
	"y": 4,
	// value 5
	"k": 5,
	// value 8
	"j": 8,
	"x": 8,
	// value 10
	"q": 10,
	"z": 10,
}

// Score returns the score of the given word
func Score(word string) int {
	score := 0
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		score += value[string(word[i])]
	}
	return score
}
