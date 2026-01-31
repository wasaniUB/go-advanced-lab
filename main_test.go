package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    int
		wantErr bool
	}{
		{name: "Factorial of 5", input: 5, want: 120, wantErr: false},
		{name: "Factorial of 10", input: 10, want: 3628800, wantErr: false},
		{name: "Factorial of 3", input: 3, want: 6, wantErr: false},
		{name: "Factorial of 1", input: 1, want: 1, wantErr: false},
		{name: "Factorial of 0", input: 0, want: 1, wantErr: false},
		{name: "Factorial of negative", input: -3, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Factorial(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Factorial(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    bool
		wantErr bool
	}{
		{name: "Prime number", input: 13, want: true, wantErr: false},
		{name: "Number 2", input: 2, want: true, wantErr: false},
		{name: "Non-prime number", input: 10, want: false, wantErr: false},
		{name: "Small Prime number", input: 7, want: true, wantErr: false},
		{name: "Number less than 2", input: 1, want: false, wantErr: true},
		{name: "Negative number", input: -5, want: false, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isPrime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("isPrime(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isPrime(%d) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exponent int
		want     int
		wantErr  bool
	}{
		{name: "Positive exponent", base: 2, exponent: 3, want: 8, wantErr: false},
		{name: "Another positive exponent", base: 5, exponent: 4, want: 625, wantErr: false},
		{name: "Zero base", base: 0, exponent: 3, want: 0, wantErr: false},
		{name: "Zero base and exponent", base: 0, exponent: 0, want: 0, wantErr: true},
		{name: "Zero exponent", base: 5, exponent: 0, want: 1, wantErr: false},
		{name: "Negative exponent", base: 2, exponent: -2, want: 0, wantErr: true},
		{name: "Another negative exponent", base: 2, exponent: -3, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Power(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Power(%d, %d) error = %v, wantErr %v", tt.base, tt.exponent, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Power(%d, %d) = %d, want %d", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}
