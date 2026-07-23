package bottlesong

import (
	"fmt"
	"strings"
)

var n = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	verses := make([]string, 0)
	for i, j := 0, startBottles; i < takeDown; i++ {
		var v1 = fmt.Sprintf("%s green bottles hanging on the wall,", n[j])
		if j == 1 {
			v1 = fmt.Sprintf("%s green bottle hanging on the wall,", n[j])
		}
		verses = append(verses, v1)
		verses = append(verses, v1)
		verses = append(verses, "And if one green bottle should accidentally fall,")
		j--
		var v2 string
		switch j {
		case 1:
			v2 = fmt.Sprintf("There'll be %s green bottle hanging on the wall.", strings.ToLower(n[j]))
		case 0:
			v2 = "There'll be no green bottles hanging on the wall."
		default:
			v2 = fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(n[j]))
		}
		verses = append(verses, v2)
		if i+1 < takeDown {
			verses = append(verses, "")
		}
	}
	return verses
}
