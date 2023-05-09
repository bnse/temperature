package temperature

import "testing"

func TestKelvin_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want string
	}{
		// TODO: Add test cases.

		{"Test String with zero Kelvin", Kelvin(0), "0 K"},
		{"Test String with positive Kelvin", Kelvin(273.15), "273.15 K"},
		{"Test String with negative Kelvin", Kelvin(-273.15), "-273.15 K"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Kelvin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_FixDecimal(t *testing.T) {
	type args struct {
		precision int
	}
	tests := []struct {
		name string
		k    Kelvin
		args args
		want Kelvin
	}{
		// TODO: Add test cases.

		{"Test FixDecimal with zero precision", Kelvin(0), args{0}, Kelvin(0)},
		{"Test FixDecimal with positive precision", Kelvin(273.156), args{2}, Kelvin(273.16)},
		{"Test FixDecimal with negative precision", Kelvin(273.156), args{-1}, Kelvin(270)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.FixDecimal(tt.args.precision); got != tt.want {
				t.Errorf("Kelvin.FixDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_AbsoluteZero(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want Kelvin
	}{
		// TODO: Add test cases.

		{"Test AbsoluteZero with zero Kelvin", Kelvin(0), Kelvin(0)},
		{"Test AbsoluteZero with positive Kelvin", Kelvin(273.15), Kelvin(0)},
		{"Test AbsoluteZero with negative Kelvin", Kelvin(-273.15), Kelvin(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsoluteZero(tt.k); got != tt.want {
				t.Errorf("Kelvin.AbsoluteZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_IsValid(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want bool
	}{
		// TODO: Add test cases.

		{"Test IsValid with zero Kelvin", Kelvin(0), true},
		{"Test IsValid with positive Kelvin", Kelvin(273.15), true},
		{"Test IsValid with negative Kelvin", Kelvin(-273.15), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.k); got != tt.want {
				t.Errorf("Kelvin.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_Celsius(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want Celsius
	}{
		// TODO: Add test cases.

		{"Test Celsius with zero Kelvin", Kelvin(0), Celsius(-273.15)},
		{"Test Celsius with positive Kelvin", Kelvin(273.15), Celsius(0)},
		{"Test Celsius with negative Kelvin", Kelvin(-273.15), Celsius(-546.3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.Celsius(); got != tt.want {
				t.Errorf("Kelvin.Celsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_Fahrenheit(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want Fahrenheit
	}{
		// TODO: Add test cases.
		{"Test Fahrenheit with zero Kelvin", Kelvin(0), Fahrenheit(-459.67)},
		{"Test Fahrenheit with positive Kelvin", Kelvin(273.15), Fahrenheit(31.999999999999943)},
		{"Test Fahrenheit with negative Kelvin", Kelvin(-273.15), Fahrenheit(-951.3399999999999)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.Fahrenheit(); got != tt.want {
				t.Errorf("Kelvin.Fahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}
