// Package bob implements a simple library that implements
// simple routine to calculate conversation responses for bob
package bob

import "strings"

// Hey returns converstaion response appropriate for remark, if successful.
func Hey(remark string) string {
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := strings.Compare(remark, strings.ToUpper(remark)) == 0
	if isYell && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isQuestion {
		return "Sure."
	}
	if isYell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
