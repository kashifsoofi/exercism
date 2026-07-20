// Package rotationalcipher implements a simple library for rotational ceasar cipher
package rotationalcipher

import (
	"unicode"
)

// RotationalCipher given plain text and shift key
// shifts all the letters and returns cipher text
func RotationalCipher(text string, shift int) string {
	shifted := make([]rune, len(text))
	for i, r := range text {
		shifted[i] = rotate(r, shift)
	}

	return string(shifted)
}

func rotate(r rune, shift int) rune {
	if !unicode.IsLetter(r) {
		return r
	}

	offset := int('a')
	if unicode.IsUpper(r) {
		offset = int('A')
	}

	s := rune(((int(r)+shift)-offset)%26 + offset)
	return s
}
