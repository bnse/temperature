package temperature

import (
	"math"
	"strconv"
)

// Kelvin represents a temperature in Kelvin.
type Kelvin float64

// Fahrenheit represents a temperature in degrees Fahrenheit.
type Fahrenheit float64

// Celsius represents a temperature in Celsius scale.
type Celsius float64

type Temperature interface {
	Celsius | Fahrenheit | Kelvin
}

// round function to round off float64 numbers
func round(num float64) int {
	return int(float64(num) + math.Copysign(0.5, float64(num)))
}

// FixDecimal function to round off float64 numbers upto a certain number of decimal places
func FixDecimal(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func AbsoluteZero[T Temperature](t T) T {
	switch any(t).(type) {
	case Celsius:
		return T(AbsoluteZeroForCelsius)
	case Fahrenheit:
		return T(AbsoluteZeroForFahrenheit)
	case Kelvin:
		return T(AbsoluteZeroForKelvin)
	}
	return T(AbsoluteZeroForKelvin)
}

func IsValid[T Temperature](t T) bool {
	return t >= AbsoluteZero(t)
}

// String returns the Celsius temperature as a string with the unit.
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', -1, 64) + " °C"
}

// String returns the Fahrenheit temperature as a string with the unit "°F".
func (f Fahrenheit) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64) + " °F"
}

func (k Kelvin) String() string {
	return strconv.FormatFloat(float64(k), 'f', -1, 64) + " K"
}

// FixDecimal returns a new Fahrenheit temperature with the specified decimal precision.
func (f Fahrenheit) FixDecimal(precision int) Fahrenheit {
	return Fahrenheit(FixDecimal(float64(f), precision))
}

const AbsoluteZeroForFahrenheit Fahrenheit = -459.67

// Celsius returns the Celsius temperature equivalent of the Fahrenheit temperature.
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Kelvin returns the Kelvin temperature equivalent of the Fahrenheit temperature.
func (f Fahrenheit) Kelvin() Kelvin {
	return Kelvin((f + 459.67) * 5 / 9)
}

const AbsoluteZeroForCelsius Celsius = -273.15

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

// FixDecimal returns a new Kelvin value with the specified number of decimal places.
func (k Kelvin) FixDecimal(precision int) Kelvin {
	return Kelvin(FixDecimal(float64(k), precision))
}

const AbsoluteZeroForKelvin Kelvin = 0

// Celsius returns the Celsius equivalent of a Kelvin temperature.
func (k Kelvin) Celsius() Celsius {
	return Celsius(k - 273.15)
}

// Fahrenheit returns the Fahrenheit equivalent of a Kelvin temperature.
func (k Kelvin) Fahrenheit() Fahrenheit {
	return Fahrenheit((k*9/5 - 459.67))
}

// TODO: https://en.wikipedia.org/wiki/Temperature
// https://en.wikipedia.org/wiki/Conversion_of_scales_of_temperature
