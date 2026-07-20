// Package protein implements a simple library that translate RNA into proteins
package protein

import (
	"errors"
)

var proteins = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "",
	"UAG": "",
	"UGA": "",
}

// ErrStop for stop codon
var ErrStop = errors.New("stop codon")

// ErrInvalidBase for invalid codon
var ErrInvalidBase = errors.New("invalid base")

// FromCodon given codon returns protein name
func FromCodon(c string) (string, error) {
	p, ok := proteins[c]
	if !ok {
		return "", ErrInvalidBase
	}

	if len(p) == 0 {
		return p, ErrStop
	}

	return p, nil
}

// FromRNA given an RNA sequence returns array of protein names
func FromRNA(rna string) ([]string, error) {
	result := []string{}
	runes := []rune(rna)
	for i := 0; i < len(runes); i += 3 {
		j := i + 3
		if j > len(runes) {
			j = len(runes)
		}
		c := string(runes[i:j])
		p, err := FromCodon(c)
		if err != nil {
			if err == ErrStop {
				return result, nil
			}

			return result, err
		}

		result = append(result, p)
	}

	return result, nil
}
