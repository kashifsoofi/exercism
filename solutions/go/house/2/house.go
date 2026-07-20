// Package house implements a simple library that
// implements methods that output
// lyrics to `This is the House that Jack Built`
package house

import "strings"

var verseLineParts = [][]string{
	{"that lay in", "house that Jack built."},
	{"that ate", "malt"},
	{"that killed", "rat"},
	{"that worried", "cat"},
	{"that tossed", "dog"},
	{"that milked", "cow with the crumpled horn"},
	{"that kissed", "maiden all forlorn"},
	{"that married", "man all tattered and torn"},
	{"that woke", "priest all shaven and shorn"},
	{"that kept", "rooster that crowed in the morn"},
	{"that belonged to", "farmer sowing his corn"},
	{"", "horse and the hound and the horn"},
}

// Verse given a verse number return specified verse of `This is the House that Jack Built`
func Verse(n int) string {
	var verse strings.Builder
	verse.WriteString("This is")
	for i := n - 1; i >= 0; i-- {
		verse.WriteString(" the ")
		verse.WriteString(verseLineParts[i][1])
		if i > 0 {
			verse.WriteString("\n")
			verse.WriteString(verseLineParts[i-1][0])
		}
	}
	return verse.String()
}

// Song returns the full text of `This is the House that Jack Built`
func Song() string {
	var song strings.Builder
	for i := 1; i <= 12; i++ {
		song.WriteString(Verse(i))
		if i < 12 {
			song.WriteString("\n\n")
		}
	}
	return song.String()
}
