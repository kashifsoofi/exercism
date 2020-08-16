// Package robotname implements a simple library that implements
// method to validate number using luhn formula
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var maximumUniqueNames = 26 * 26 * 10 * 10 * 10
var usedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Robot struct to keep robot name
type Robot struct {
	name string
}

// Name generates random name if not set
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}

	if len(usedNames) == maximumUniqueNames {
		return "", errors.New("no more unique names available")
	}

	name := generateRandomName()
	for usedNames[name] {
		name = generateRandomName()
	}

	usedNames[name] = true
	r.name = name
	return r.name, nil
}

// Reset resets Robot name
func (r *Robot) Reset() {
	r.name = ""
}

func generateRandomName() string {
	c1 := rune('A' + rand.Intn(26))
	c2 := rune('A' + rand.Intn(26))
	number := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", c1, c2, number)
}
