package temperature

import "testing"

func TestFixDecimal(t *testing.T) {
	type args struct {
		num       float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test case 1", args{3.14159, 2}, 3.14},
		{"Test case 2", args{2.71828, 3}, 2.718},
		{"Test case 3", args{1.23456789, 5}, 1.23457},
		{"Test case 4", args{0.123456789, 8}, 0.12345679},
		{"Test case 5", args{100.0, 0}, 100.0},
		{"Test case 6", args{0.0, 2}, 0.0},
		{"Test case 7", args{-3.14159, 2}, -3.14},
		{"Test case 8", args{-2.71828, 3}, -2.718},
		{"Test case 9", args{-1.23456789, 5}, -1.23457},
		{"Test case 10", args{-0.123456789, 8}, -0.12345679},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixDecimal(tt.args.num, tt.args.precision); got != tt.want {
				t.Errorf("FixDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
