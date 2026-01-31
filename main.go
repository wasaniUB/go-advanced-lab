package main

import "fmt"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("prime check requires number >= 2")
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, fmt.Errorf("negative exponents not supported")
	}
	if exponent == 0 && base == 0 {
		return 0, fmt.Errorf("MATH ERROR!! 0^0 is undefined")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}
