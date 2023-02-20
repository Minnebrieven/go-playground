package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(count int) string {
	switch {
	case count%3 == 0 && count%5 == 0:
		return "FizzBuzz"
	case count%3 == 0:
		return "Fizz"
	case count%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(count)
	}

}

func main() {
	n := 1
	for n <= 100 {
		fmt.Printf("%v, ", fizzBuzz(n))
		n++
	}
	fmt.Println("")
}
