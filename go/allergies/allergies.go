// Package allergies implements a simple library that
// implements methods to determine if a person is allergic
// to a given item and provide full list of person's allergies
package allergies

const (
	eggsIndex         uint = iota // eggs index
	peanutsIndex                  // peanuts index
	shellfishIndex                // shellfish index
	strawberriesIndex             // strawberries index
	tomatoesIndex                 // reverse index
	chocolateIndex                // chocolate index
	pollenIndex                   // pollen index
	catsIndex                     // cats index
)

const (
	eggs         uint = 1 << iota // eggs
	peanuts                       // peanuts
	shellfish                     // shellfish
	strawberries                  // strawberries
	tomatoes                      // reverse
	chocolate                     // chocolate
	pollen                        // pollen
	cats                          // cats
)

// Need an orderd list to get list of allergies
var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

var allergenCode = map[string]uint{
	allergens[eggsIndex]:         eggs,
	allergens[peanutsIndex]:      peanuts,
	allergens[shellfishIndex]:    shellfish,
	allergens[strawberriesIndex]: strawberries,
	allergens[tomatoesIndex]:     tomatoes,
	allergens[chocolateIndex]:    chocolate,
	allergens[pollenIndex]:       pollen,
	allergens[catsIndex]:         cats,
}

// Allergies given score, return a list of allergies
func Allergies(score uint) []string {
	allergies := make([]string, 0)
	for _, a := range allergens {
		if score&allergenCode[a] == allergenCode[a] {
			allergies = append(allergies, a)
		}
	}
	return allergies
}

// AllergicTo given score and item, returns true if person is allergic to that item
func AllergicTo(score uint, item string) bool {
	return score&allergenCode[item] == allergenCode[item]
}
