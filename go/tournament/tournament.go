// Package tournament implements a simple library that implements
// method to tally the results of a small football competition
package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type teamStat struct {
	name   string
	wins   int
	losses int
	draws  int
}

func (t *teamStat) addWin() {
	t.wins++
}

func (t *teamStat) addLoss() {
	t.losses++
}

func (t *teamStat) addDraw() {
	t.draws++
}

func (t *teamStat) matchesPlayed() int {
	return t.wins + t.losses + t.draws
}

func (t *teamStat) points() int {
	return t.wins*3 + t.draws
}

// Tally reads match results from reader, converts it to table format and writes to writer
func Tally(reader io.Reader, writer io.Writer) error {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("error reading input")
	}

	resultLines := strings.Split(string(input), "\n")
	stats, err := parseResults(resultLines)
	if err != nil {
		return err
	}

	sort.Slice(stats, func(i, j int) bool {
		if stats[i].points() != stats[j].points() {
			return stats[i].points() > stats[j].points()
		}
		return stats[i].name < stats[j].name
	})

	writer.Write([]byte(fmt.Sprintf("%-30v | MP |  W |  D |  L |  P\n", "Team")))
	for _, s := range stats {
		writer.Write([]byte(fmt.Sprintf("%-30v | %2d | %2d | %2d | %2d | %2d\n", s.name, s.matchesPlayed(), s.wins, s.draws, s.losses, s.points())))
	}

	return nil
}

func parseResults(results []string) ([]*teamStat, error) {
	m := make(map[string]*teamStat)

	for _, r := range results {
		if len(r) == 0 || strings.HasPrefix(r, "#") {
			continue
		}

		values := strings.Split(r, ";")
		if len(values) != 3 {
			return nil, errors.New("invalid result")
		}

		s1 := m[values[0]]
		if s1 == nil {
			s1 = &teamStat{
				name: values[0],
			}
			m[s1.name] = s1
		}

		s2 := m[values[1]]
		if s2 == nil {
			s2 = &teamStat{
				name: values[1],
			}
			m[s2.name] = s2
		}

		switch values[2] {
		case "win":
			s1.addWin()
			s2.addLoss()
			break
		case "loss":
			s1.addLoss()
			s2.addWin()
			break
		case "draw":
			s1.addDraw()
			s2.addDraw()
			break
		default:
			return nil, errors.New("invalid result")
		}
	}

	stats := make([]*teamStat, 0, len(m))
	for _, v := range m {
		stats = append(stats, v)
	}
	return stats, nil
}
