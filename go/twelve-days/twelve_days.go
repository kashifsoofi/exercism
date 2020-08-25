// Package twelve implements a simple library that
// implements methods that output
// lyrics to `The Twelve Days of Christmas`
package twelve

import (
	"fmt"
	"strings"
)

var dayName = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Verse given a verse number return specified verse from `The Twelve Days of Christmas`
func Verse(n int) string {
	var giftsFoDay strings.Builder
	for i := n - 1; i > 0; i-- {
		giftsFoDay.WriteString(gifts[i])
		giftsFoDay.WriteString(", ")
	}
	if giftsFoDay.Len() > 0 {
		giftsFoDay.WriteString("and ")
	}
	giftsFoDay.WriteString(gifts[0])

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", dayName[n], giftsFoDay.String())
}

// Song returns complete `The Twelve Days of Christmas` song
func Song() string {
	var song strings.Builder
	for n := 1; n <= 12; n++ {
		song.WriteString(Verse(n))
		if n < 12 {
			song.WriteString("\n")
		}
	}
	return song.String()
}
