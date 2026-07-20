// Package kindergarten implements utility methods to determine children responsible for each plant
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plants = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// Garden represent Kingergarten Garden
type Garden map[string][]string

// NewGarden given diagram and list of children creates and returns a new instance of Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	if !strings.HasPrefix(diagram, "\n") {
		return nil, errors.New("invalid diagram")
	}

	rows := strings.Split(diagram[1:], "\n")
	if len(rows) != 2 ||
		len(rows[0]) != len(rows[1]) ||
		len(rows[0])%2 != 0 ||
		len(rows[0]) != len(children)*2 {
		return nil, errors.New("invalid diagram")
	}

	sorted := append([]string{}, children...)
	sort.Strings(sorted)
	for i := 0; i+1 < len(sorted); i++ {
		if sorted[i] == sorted[i+1] {
			return nil, errors.New("duplicate child name")
		}
	}

	g := Garden{}
	for _, row := range rows {
		for i, r := range row {
			if _, ok := plants[r]; !ok {
				return nil, errors.New("invalid plant code")
			}

			child := sorted[i/2]
			g[child] = append(g[child], plants[r])
		}
	}
	return &g, nil
}

// Plants returns list of plants child is responsible for
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
