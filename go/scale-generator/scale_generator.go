// Package scale implements utility routine to generate scales
package scale

import (
	"strings"
)

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var useSharps = map[string]bool{
	"G":  true,
	"D":  true,
	"A":  true,
	"E":  true,
	"B":  true,
	"F#": true,
	"e":  true,
	"b":  true,
	"f#": true,
	"c#": true,
	"g#": true,
	"d#": true,
	"C":  true,
	"a":  true,
}

var intervalMap = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// Scale returns scale for given tonic and interval
func Scale(tonic, interval string) []string {
	p := []string{}
	if useSharps[tonic] == true {
		p = sharps
	} else {
		p = flats
	}

	if len(interval) == 0 {
		interval = "mmmmmmmmmmmm"
	}

	pi := 0
	tonic = strings.ToTitle(tonic[:1]) + tonic[1:]
	for i, v := range p {
		if v == tonic {
			pi = i
			break
		}
	}

	s := []string{}
	for _, i := range interval {
		s = append(s, p[pi])
		pi = (pi + intervalMap[i]) % len(p)
	}

	return s
}
