// Package robotname implements a simple library that implements
// method to validate number using luhn formula
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var maximumUniqueNames = 26 * 26 * 10 * 10 * 10

type Robot struct {
	name string
}

var generatedNames = make(map[string]bool)

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := generateName()
		if err != nil {
			return "", err
		}

		r.name = name
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	// delete(generatedNames, r.name)
	r.name = ""
}

func generateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	i1 := rand.Intn(26)
	i2 := rand.Intn(26)
	number := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", alphabet[i1], alphabet[i2], number)
}

func generateName() (string, error) {
	if len(generatedNames) == maximumUniqueNames {
		return "", errors.New("no more unique names available")
	}

	name := generateRandomName()
	for generatedNames[name] {
		name = generateRandomName()
	}

	generatedNames[name] = true
	return name, nil
}
