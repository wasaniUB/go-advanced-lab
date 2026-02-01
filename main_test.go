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

func TestMakeCounter(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		calls    int
		expected []int
	}{
		{name: "Start at 0", start: 0, calls: 5, expected: []int{1, 2, 3, 4, 5}},
		{name: "Start at 10", start: 10, calls: 3, expected: []int{11, 12, 13}},
		{name: "Start at -5", start: -5, calls: 4, expected: []int{-4, -3, -2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := makeCounter(tt.start)
			for i := 0; i < tt.calls; i++ {
				got := counter()
				if got != tt.expected[i] {
					t.Errorf("Call %d: got %d, want %d", i+1, got, tt.expected[i])
				}
			}
		})
	}

	counter1 := makeCounter(0)
	counter2 := makeCounter(100)

	if counter1() != 1 {
		t.Errorf("counter1 first call = %d, want 1", counter1())
	}
	if counter2() != 101 {
		t.Errorf("counter2 first call = %d, want 101", counter2())
	}
	if counter1() != 2 {
		t.Errorf("counter1 second call = %d, want 2", counter1())
	}
}

func TestMakeMultiplier(t *testing.T) {
	tests := []struct {
		name     string
		factor   int
		inputs   []int
		expected []int
	}{
		{name: "Double", factor: 2, inputs: []int{5}, expected: []int{10}},
		{name: "Triple", factor: 3, inputs: []int{5}, expected: []int{15}},
		{name: "Multiply by zero", factor: 0, inputs: []int{10}, expected: []int{0}},
		{name: "Negative factor", factor: -2, inputs: []int{4}, expected: []int{-8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := MakeMultiplier(tt.factor)
			for i, input := range tt.inputs {
				got := multiplier(input)
				if got != tt.expected[i] {
					t.Errorf("Input %d: got %d, want %d", input, got, tt.expected[i])
				}
			}
		})
	}
}

func TestMakeAccumulator(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		addValue int
		subValue int
		expected int
	}{
		{name: "Add then subtract", initial: 100, addValue: 50, subValue: 30, expected: 120},
		{name: "Adding only", initial: 10, addValue: 15, subValue: 0, expected: 25},
		{name: "Subtracting only", initial: 20, addValue: 0, subValue: 5, expected: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add, subtract, get := makeAccumulator(tt.initial)

			if tt.addValue != 0 {
				add(tt.addValue)
			}
			if tt.subValue != 0 {
				subtract(tt.subValue)
			}

			got := get()
			if got != tt.expected {
				t.Errorf("final value = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		operation func(int) int
		expected  []int
	}{
		{name: "Square numbers", input: []int{1, 2, 3, 4, 100}, operation: func(x int) int { return x * x }, expected: []int{1, 4, 9, 16, 10000}},
		{name: "Double numbers", input: []int{1, 2, 3, 2008, 2002}, operation: func(x int) int { return x * 2 }, expected: []int{2, 4, 6, 4016, 4004}},
		{name: "Negate numbers", input: []int{1, -2, 3}, operation: func(x int) int { return -x }, expected: []int{-1, 2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]int{}, tt.input...)
			got := Apply(tt.input, tt.operation)

			if len(got) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], tt.expected[i])
				}
			}

			for i := range tt.input {
				if tt.input[i] != original[i] {
					t.Errorf("original slice was modified")
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{name: "Even numbers", input: []int{1, 2, 3, 4, 5, 6}, predicate: func(x int) bool { return x%2 == 0 }, expected: []int{2, 4, 6}},
		{name: "Positive numbers", input: []int{-3, -1, 0, 2, 5}, predicate: func(x int) bool { return x > 0 }, expected: []int{2, 5}},
		{name: "Numbers greater than 10", input: []int{5, 10, 15, 20}, predicate: func(x int) bool { return x > 10 }, expected: []int{15, 20}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.input, tt.predicate)

			if len(got) != len(tt.expected) {
				t.Errorf("length mismatch: got %d, want %d", len(got), len(tt.expected))
			}

			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		initial   int
		operation func(int, int) int
		expected  int
	}{
		{name: "Sum", input: []int{1, 2, 3, 4}, initial: 0, operation: func(acc, curr int) int { return acc + curr }, expected: 10},
		{name: "Product", input: []int{1, 2, 3, 4}, initial: 1, operation: func(acc, curr int) int { return acc * curr }, expected: 24},
		{name: "Maximum", input: []int{3, 1, 4, 2}, initial: 0, operation: func(acc, curr int) int {
			if curr > acc {
				return curr
			}
			return acc
		},
			expected: 4,
		},
		{name: "Minimum", input: []int{3, 1, 4, 2}, initial: 100, operation: func(acc, curr int) int {
			if curr < acc {
				return curr
			}
			return acc
		},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reduce(tt.input, tt.initial, tt.operation)
			if got != tt.expected {
				t.Errorf("got %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	tests := []struct {
		name     string
		f        func(int) int
		g        func(int) int
		input    int
		expected int
	}{
		{name: "Double then add two", f: func(x int) int { return x + 2 }, g: func(x int) int { return x * 2 }, input: 5, expected: 12},
		{name: "Square then negate", f: func(x int) int { return -x }, g: func(x int) int { return x * x }, input: 4, expected: -16},
		{name: "Add three then multiply by four", f: func(x int) int { return x * 4 }, g: func(x int) int { return x + 3 }, input: 2, expected: 20},
		{name: "Identity functions", f: func(x int) int { return x }, g: func(x int) int { return x }, input: 10, expected: 10},
		{name: "Cube then square", f: func(x int) int { return x * x }, g: func(x int) int { return x * x * x }, input: 2, expected: 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := Compose(tt.f, tt.g)
			got := composed(tt.input)

			if got != tt.expected {
				t.Errorf("got %d, want %d", got, tt.expected)
			}
		})
	}
}
