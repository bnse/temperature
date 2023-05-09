package temperature

import (
	"strconv"
)

// Fahrenheit represents a temperature in degrees Fahrenheit.
type Fahrenheit float64

// String returns the Fahrenheit temperature as a string with the unit "°F".
func (f Fahrenheit) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64) + " °F"
}

// FixDecimal returns a new Fahrenheit temperature with the specified decimal precision.
func (f Fahrenheit) FixDecimal(precision int) Fahrenheit {
	return Fahrenheit(FixDecimal(float64(f), precision))
}

const AbsoluteZeroForFahrenheit Fahrenheit = -459.67

// AbsoluteZero returns the Fahrenheit temperature of absolute zero (-459.67 °F).
func (f Fahrenheit) AbsoluteZero() Fahrenheit {
	return Fahrenheit(AbsoluteZeroForFahrenheit)
}

// IsValid returns true if the Fahrenheit temperature is valid (above absolute zero).
func (f Fahrenheit) IsValid() bool {
	return f >= f.AbsoluteZero()
}

// Celsius returns the Celsius temperature equivalent of the Fahrenheit temperature.
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Kelvin returns the Kelvin temperature equivalent of the Fahrenheit temperature.
func (f Fahrenheit) Kelvin() Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}
