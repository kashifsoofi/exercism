// Package beer implements simple routings that output lyrics to `99 Bottles of Beer on the Wall`
package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns a verse from `99 Bottles of Beer on the Wall`
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("invalid verse")
	}

	v := ""
	switch n {
	case 0:
		v = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	case 1:
		v = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	case 2:
		v = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
	default:
		v = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
	}

	return v, nil
}

// Verses returns numbered verses from `99 Bottles of Beer on the Wall`
func Verses(upper, lower int) (string, error) {
	if upper < lower {
		return "", errors.New("invalid bounds")
	}

	var verses strings.Builder
	for i := upper; i >= lower; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses.WriteString(v)
		verses.WriteString("\n")
	}
	return verses.String(), nil
}

// Song returns complete `99 Bottles of Beer on the Wall` song
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
