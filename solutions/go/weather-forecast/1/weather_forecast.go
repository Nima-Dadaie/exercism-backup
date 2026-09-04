// Package weather provides tools to forecast the weather condition of cities.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents the current location being forecasted.
	CurrentLocation  string
)
// Forecast returns the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
