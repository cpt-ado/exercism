// Package space calculates the age of a user
// on any planet in the Solar System
package space

type Planet string

var ratio = map[Planet]float64{
    "Mercury": 0.2408467,
    "Venus": 0.61519726,
    "Earth": 1.0,
    "Mars": 1.8808158,
    "Jupiter": 11.862615,
    "Saturn": 29.447498,
    "Uranus": 84.016846,
    "Neptune": 164.79132,
}

// Age returns the number of cycles on the
// given planet for the given seconds
func Age(seconds float64, planet Planet) float64 {
    planetYear := ratio[planet] * 31557600 // seconds in year on Earth
    return seconds / planetYear
}
