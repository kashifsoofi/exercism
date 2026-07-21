package bottlesong

import (
	"fmt"
	"strings"
)

var englishNumbers = []string{
	"No",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
}

func Recite(startBottles, takeDown int) []string {
	var verses []string
	for i := 0; i < takeDown; i++ {
		verses = append(verses, verse(startBottles)...)
		if i+1 < takeDown {
			verses = append(verses, "")
		}
		startBottles--
	}
	return verses
}

func verse(startBottles int) []string {
	line1 := fmt.Sprintf("%s green bottle%s hanging on the wall,", englishNumbers[startBottles], plural(startBottles))
	line3 := "And if one green bottle should accidentally fall,"
	line4 := fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", strings.ToLower(englishNumbers[startBottles-1]), plural(startBottles-1))
	return []string{line1, line1, line3, line4}
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
