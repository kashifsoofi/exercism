// Package proverb implements a utility method to return proverbs given rhymes
package proverb

// Proverb returns list of proverbs given list of rhymes
func Proverb(rhyme []string) []string {
	proverbs := make([]string, len(rhyme))
	if len(rhyme) == 0 {
		return proverbs
	}

	for i := 0; i < len(rhyme)-1; i++ {
		proverbs[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	proverbs[len(proverbs)-1] = "And all for the want of a " + rhyme[0] + "."
	return proverbs
}
