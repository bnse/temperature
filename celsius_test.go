package temperature

import (
	"testing"
)

func TestCelsius_String(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
		// TODO: Add test cases.
		{"zero celsius", Celsius(0), "0 °C"},
		{"positive celsius", Celsius(25), "25 °C"},
		{"negative celsius", Celsius(-10), "-10 °C"},
		{"decimal celsius", Celsius(12.3456), "12.3456 °C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Celsius.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCelsius_FixDecimal(t *testing.T) {
	type args struct {
		precision int
	}
	tests := []struct {
		name string
		c    Celsius
		args args
		want Celsius
	}{
		// TODO: Add test cases.
		{"zero celsius", Celsius(0), args{2}, Celsius(0)},
		{"positive celsius", Celsius(25), args{2}, Celsius(25)},
		{"negative celsius", Celsius(-10), args{2}, Celsius(-10)},
		{"decimal celsius", Celsius(12.3456), args{2}, Celsius(12.35)},
		{"decimal celsius", Celsius(12.3456), args{3}, Celsius(12.346)},
		{"decimal celsius", Celsius(12.3456), args{4}, Celsius(12.3456)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.FixDecimal(tt.args.precision); got != tt.want {
				t.Errorf("Celsius.FixDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCelsius_Fahrenheit(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
		// TODO: Add test cases.
		{"zero celsius", Celsius(0), Fahrenheit(32)},
		{"positive celsius", Celsius(25), Fahrenheit(77)},
		{"negative celsius", Celsius(-10), Fahrenheit(14)},
		{"decimal celsius", Celsius(12.3456), Fahrenheit(54.22208)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Fahrenheit(); got != tt.want {
				t.Errorf("Celsius.Fahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCelsius_Kelvin(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Kelvin
	}{
		// TODO: Add test cases.
		{"zero celsius", Celsius(0), Kelvin(273.15)},
		{"positive celsius", Celsius(25), Kelvin(298.15)},
		{"negative celsius", Celsius(-10), Kelvin(263.15)},
		{"decimal celsius", Celsius(12.3456), Kelvin(285.49559999999997)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Kelvin(); got != tt.want {
				t.Errorf("Celsius.Kelvin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCelsius_AbsoluteZero(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Celsius
	}{
		// TODO: Add test cases.
		{"absolute zero", Celsius(-273.15), Celsius(-273.15)},
		{"above absolute zero", Celsius(-273.14), Celsius(-273.15)},
		{"below absolute zero", Celsius(-273.16), Celsius(-273.15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AbsoluteZero(); got != tt.want {
				t.Errorf("Celsius.AbsoluteZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
