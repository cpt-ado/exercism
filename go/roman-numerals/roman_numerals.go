// Package romannumerals gets the roman representation of given integers.
package romannumerals

import (
	"errors"
	"strings"
)

// Constant steps to use identifying roman numerals.
func romanSteps() []int {
	return []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
}

// Constant representation of integers to roman numerals.
func romanLetter() map[int]string {
	return map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
}

// Returns a the roman representation of any integer between 1 and 3000.
func ToRomanNumeral(arabic int) (string, error) {
	roman := ""
	if arabic < 1 || arabic > 3000 {
		return roman, errors.New("input from invalid range")
	}
	for _, step := range romanSteps() {
		quotient := arabic / step
		if quotient > 0 {
			roman += strings.Repeat(
				romanLetter()[step], quotient)
			arabic = arabic % step
		}
	}
	return roman, nil
}
