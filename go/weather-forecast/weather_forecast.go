// Package weather returns weather.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current city or location.
var CurrentLocation string

// Forecast function returns CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
