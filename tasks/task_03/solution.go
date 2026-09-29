package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	// Проверка входных данных
	if n == 0 {
		return "", ErrZero
	}
	
	var out string
	if n % 3 == 0 {
		out += "Fizz"
	}
	if n % 5 == 0 {
		out += "Buzz"
	}
	
	if out == "" {
		return strconv.Itoa(n), nil
	}
	return out, nil
}
