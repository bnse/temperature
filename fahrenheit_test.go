package temperature

import "testing"

func TestFahrenheit_String(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		// TODO: Add test cases.
		{"positive value", Fahrenheit(32), "32 °F"},
		{"negative value", Fahrenheit(-10), "-10 °F"},
		{"zero value", Fahrenheit(0), "0 °F"},
		{"decimal value", Fahrenheit(98.6), "98.6 °F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Fahrenheit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_FixDecimal(t *testing.T) {
	type args struct {
		precision int
	}
	tests := []struct {
		name string
		f    Fahrenheit
		args args
		want Fahrenheit
	}{
		// TODO: Add test cases.

		{"no decimal", Fahrenheit(32), args{0}, Fahrenheit(32)},
		{"one decimal", Fahrenheit(98.6), args{1}, Fahrenheit(98.6)},
		{"two decimal", Fahrenheit(98.654), args{2}, Fahrenheit(98.65)},
		{"three decimal", Fahrenheit(98.6543), args{3}, Fahrenheit(98.654)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FixDecimal(tt.args.precision); got != tt.want {
				t.Errorf("Fahrenheit.FixDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_AbsoluteZero(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want Fahrenheit
	}{
		// TODO: Add test cases.
		{"absolute zero", Fahrenheit(-459.67), Fahrenheit(-459.67)},
		{"above absolute zero", Fahrenheit(-459.66), Fahrenheit(-459.67)},
		{"below absolute zero", Fahrenheit(-459.68), Fahrenheit(-459.67)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.AbsoluteZero(); got != tt.want {
				t.Errorf("Fahrenheit.AbsoluteZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_IsValid(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want bool
	}{
		// TODO: Add test cases.
		{"valid temperature", Fahrenheit(32), true},
		{"valid temperature", Fahrenheit(-10), true},
		{"valid temperature", Fahrenheit(0), true},
		{"valid temperature", Fahrenheit(98.6), true},
		{"invalid temperature", Fahrenheit(-500), false},
		{"invalid temperature", Fahrenheit(-459.68), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsValid(); got != tt.want {
				t.Errorf("Fahrenheit.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_Celsius(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want Celsius
	}{
		// TODO: Add test cases.

		{"freezing point", Fahrenheit(32), Celsius(0)},
		{"boiling point", Fahrenheit(212), Celsius(100)},
		{"body temperature", Fahrenheit(98.6), Celsius(37)},
		{"absolute zero", Fahrenheit(-459.67), Celsius(-273.15)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Celsius(); got != tt.want {
				t.Errorf("Fahrenheit.Celsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_Kelvin(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want Kelvin
	}{
		// TODO: Add test cases.

		{"freezing point", Fahrenheit(32), Kelvin(273.15)},
		{"boiling point", Fahrenheit(212), Kelvin(373.15000000000003)},
		{"body temperature", Fahrenheit(98.6), Kelvin(310.15)},
		{"absolute zero", Fahrenheit(-459.67), Kelvin(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Kelvin(); got != tt.want {
				t.Errorf("Fahrenheit.Kelvin() = %v, want %v", got, tt.want)
			}
		})
	}
}
