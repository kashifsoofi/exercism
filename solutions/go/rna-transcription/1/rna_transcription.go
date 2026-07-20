// Package strand implements a simple library for rna transcription
package strand

import "strings"

var compliments = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA given a dna string returns its RNA compliment
func ToRNA(dna string) string {
	var rna strings.Builder
	for _, r := range dna {
		rna.WriteRune(compliments[r])
	}
	return rna.String()
}
