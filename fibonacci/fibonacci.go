package fibonacci

import (
	"errors"
)

var (
	ErrToCannotBeZeroOrLess     = errors.New("numbers to cannot be zero or less")
	ErrFromCannotBeLessThanZero = errors.New("numbers from cannot be less than zero")
	ErrFromGreaterThanTo        = errors.New("from is greater than to")
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func FibonacciCalculation(from, to int) ([]int, error) {
	if from < 0 {
		return []int{}, ErrFromCannotBeLessThanZero
	}
	if to <= 0 {
		return []int{}, ErrToCannotBeZeroOrLess
	}
	if to < from {
		return []int{}, ErrFromGreaterThanTo
	}
	var length = to - from + 1
	var result = make([]int, 0, length)
	result = append(result, fibonacci(from), fibonacci(from+1))
	for i := 2; i < length; i++ {
		result = append(result, result[i-2]+result[i-1])
	}
	return result, nil
}
