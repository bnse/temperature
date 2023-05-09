package temperature

import (
	"math"
)

// round function to round off float64 numbers
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// FixDecimal function to round off float64 numbers upto a certain number of decimal places
func FixDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// TODO: https://en.wikipedia.org/wiki/Temperature
// https://en.wikipedia.org/wiki/Conversion_of_scales_of_temperature
