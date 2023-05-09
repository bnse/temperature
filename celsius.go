package temperature

import (
	"strconv"
)

// Celsius represents a temperature in Celsius scale.
type Celsius float64

const AbsoluteZeroForCelsius Celsius = -273.15

// AbsoluteZero returns the Celsius temperature of absolute zero (-273.15 °C).
func (c Celsius) AbsoluteZero() Celsius {
	return AbsoluteZeroForCelsius
}

// IsValid returns true if the Celsius temperature is valid (above absolute zero).
func (c Celsius) IsValid() bool {
	return c >= c.AbsoluteZero()
}

// String returns the Celsius temperature as a string with the unit.
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', -1, 64) + " °C"
}

// FixDecimal returns a new Celsius temperature with the specified decimal precision.
func (c Celsius) FixDecimal(precision int) Celsius {
	return Celsius(FixDecimal(float64(c), precision))
}

// Fahrenheit returns the equivalent Fahrenheit temperature.
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Kelvin returns the equivalent Kelvin temperature.
func (c Celsius) Kelvin() Kelvin {
	return Kelvin(c + 273.15)
}
