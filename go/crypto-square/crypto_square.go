// Package cryptosquare implements a simple library
// for composing secret messages using method called a square code.
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encodes plain text using square code and returns cipher text.
func Encode(pt string) string {
	nt := normalise(pt)

	rows, cols := calculateColumnsAndRows(len(nt))

	var ct strings.Builder
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			ri := i + j*cols
			if ri < len(nt) {
				ct.WriteRune(rune(nt[ri]))
			} else if i < cols {
				ct.WriteRune(' ')
			}
		}
		if i+1 < cols {
			ct.WriteRune(' ')
		}
	}

	return ct.String()
}

func normalise(pt string) string {
	var nt strings.Builder
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			nt.WriteRune(unicode.ToLower(r))
		}
	}
	// The following is not working
	// nt = strings.TrimFunc(pt, func(r rune) bool {
	// 	return !(unicode.IsLetter(r) || unicode.IsNumber(r))
	// })
	return nt.String()
}

func calculateColumnsAndRows(l int) (int, int) {
	rows := int(math.Sqrt(float64(l)))
	cols := rows

	if rows*cols < l {
		cols++
	}

	if rows*cols < l {
		rows++
	}

	return rows, cols
}
