// Package romannumerals implements simple routine to convert numbers to roman numerals
package romannumerals

import (
	"errors"
	"sort"
	"strings"
)

var symbols = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// ToRomanNumeral returns roman numeral for a given number
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("invalid number")
	}

	keys := make([]int, len(symbols))
	i := 0
	for k := range symbols {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	v := n
	var numerals strings.Builder
	for _, k := range keys {
		for v >= k {
			numerals.WriteString(symbols[k])
			v -= k
		}

		if v == 0 {
			break
		}
	}

	return numerals.String(), nil
}
