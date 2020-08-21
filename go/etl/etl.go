// Package etl implements a simple library that
// implements transform step of an Extract-Transform-Load
package etl

import (
	"strings"
)

// Transform given Scrabble score data in an old format
// transforms and returns data in new format
func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}
	for k, v := range input {
		for _, s := range v {
			output[strings.ToLower(s)] = k
		}
	}
	return output
}
