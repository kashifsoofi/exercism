// Package space implements a simple library that
// implements methods to calculate age on different planets
package space

// Planet name
type Planet string

const secondsInEarthYear = 31557600.0

var planetSecondsInYear = map[Planet]float64{
	"Mercury": secondsInEarthYear * 0.2408467,
	"Venus":   secondsInEarthYear * 0.61519726,
	"Earth":   secondsInEarthYear,
	"Mars":    secondsInEarthYear * 1.8808158,
	"Jupiter": secondsInEarthYear * 11.862615,
	"Saturn":  secondsInEarthYear * 29.447498,
	"Uranus":  secondsInEarthYear * 84.016846,
	"Neptune": secondsInEarthYear * 164.79132,
}

const secondsOnEarth float64 = 31557600.0

// Age calculates and returns age on given planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetSecondsInYear[planet]
}
