package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	if (len(a) != len(b)) {
		return -1, errors.New("Sequences are not of equal length")
	}

	var distance = 0;
	for i, aCell := range a {
		bCell, _ := utf8.DecodeRuneInString(b[i:])
		if aCell != bCell {
			distance += 1
		}
	}

	return distance, nil
}
