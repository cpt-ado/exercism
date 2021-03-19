// Package proverb builds a proverb rhyme from given words
package proverb

// Proverb returns a proverb made from the given array of strings
func Proverb(rhyme []string) []string {
	proverb := []string{}
	length := len(rhyme)
	if length == 0 {
		return proverb
	} else if length > 1 {
		for i := 1; i < length; i++ {
			proverb = append(proverb, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
		}
	}
	return append(proverb, "And all for the want of a "+rhyme[0]+".")
}
