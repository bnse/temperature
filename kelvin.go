package temperature

import (
	"strconv"
)

// Kelvin represents a temperature in Kelvin.
type Kelvin float64

func (k Kelvin) String() string {
	return strconv.FormatFloat(float64(k), 'f', -1, 64) + " K"
}

// FixDecimal returns a new Kelvin value with the specified number of decimal places.
func (k Kelvin) FixDecimal(precision int) Kelvin {
	return Kelvin(FixDecimal(float64(k), precision))
}

const AbsoluteZeroForKelvin Kelvin = 0

// AbsoluteZero is the lowest possible temperature in Kelvin.
func (k Kelvin) AbsoluteZero() Kelvin {
	return AbsoluteZeroForKelvin
}

// IsValid returns true if the Kelvin temperature is valid (i.e., greater than or equal to AbsoluteZero).
func (k Kelvin) IsValid() bool {
	return k >= k.AbsoluteZero()
}

// Celsius returns the Celsius equivalent of a Kelvin temperature.
func (k Kelvin) Celsius() Celsius {
	return Celsius(k - 273.15)
}

// Fahrenheit returns the Fahrenheit equivalent of a Kelvin temperature.
func (k Kelvin) Fahrenheit() Fahrenheit {
	return Fahrenheit((k*9/5 - 459.67))
}
