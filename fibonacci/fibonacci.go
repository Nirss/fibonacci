package fibonacci

import (
	"errors"
)

var (
	ErrFromOrToCannotBeZeroOrLess = errors.New("numbers from or to cannot be zero or less")
	ErrFromGreaterThanTo          = errors.New("from is greater than to")
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func RecFibonacciCalculation(from, to int) ([]int, error) {
	if from <= 0 || to <= 0 {
		return []int{}, ErrFromOrToCannotBeZeroOrLess
	}
	if to < from {
		return []int{}, ErrFromGreaterThanTo
	}
	var length = to - from + 1
	var result = make([]int, 0, length)
	result = append(result, fibonacci(from-1), fibonacci(from))
	for i := 2; i < length; i++ {
		result = append(result, result[i-2]+result[i-1])
	}
	return result, nil
}
