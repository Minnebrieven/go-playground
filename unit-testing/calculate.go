package unit_testing

import (
	"errors"
)

func Addition(number, number2 int) int {
	return number + number2
}

func Subtraction(number, number2 int) int {
	return number - number2
}

func Division(number, number2 int) (int, error) {
	if number2 == 0 {
		return -1, errors.New("Number2 cant be 0")
	}
	return number / number2, nil
}

func Multiplication(number, number2 int) int {
	return number * number2
}
