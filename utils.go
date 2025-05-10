package main

import (
	"errors"
)

// Factorial calculates the factorial of a non-negative integer n.
// It returns 0 as the factorial value if n is negative.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Factoriel is not defined for negative numbers")
	}

	if n == 0 {
		return 1, nil // Factorial of 0 is 1
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial, nil
}
