// Package beer implements simple routings that output lyrics to `99 Bottles of Beer on the Wall`
package beer

import (
	"errors"
	"fmt"
	"strings"
)

var verses = [4]string{
	"No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n",
	"1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n",
	"2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n",
	"%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n",
}

// Verse returns a verse from `99 Bottles of Beer on the Wall`
func Verse(n int) (v string, err error) {
	switch {
	case n == 0, n == 1, n == 2:
		v = verses[n]
	case n > 2 && n < 100:
		v = fmt.Sprintf(verses[3], n, n, n-1)
	default:
		v, err = "", errors.New("invalid verse")
	}
	return
}

// Verses returns numbered verses from `99 Bottles of Beer on the Wall`
func Verses(upper, lower int) (string, error) {
	if upper <= lower {
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
