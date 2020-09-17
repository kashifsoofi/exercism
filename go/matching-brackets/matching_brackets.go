// Package brackets implements utility routine to match openening and closing brackets
package brackets

var openingBrackets = map[rune]bool{
	'[': true,
	'(': true,
	'{': true,
}

var matchingBrackets = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
}

// Bracket return true if opening and closing brackets are paired
func Bracket(input string) bool {
	s := []rune{}

	for _, r := range input {
		if _, ok := openingBrackets[r]; ok {
			s = append(s, r)
		} else if _, ok := matchingBrackets[r]; ok {
			if len(s) == 0 {
				return false
			}

			openingBracket := s[len(s)-1]
			if openingBracket != matchingBrackets[r] {
				return false
			}
			s = s[:len(s)-1]
		}
	}

	return len(s) == 0
}
