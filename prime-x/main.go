package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	var primeCounter, currentNumber int = 0, 2

	for {
		primeBool := func(num int) bool {
			isPrime := true
			for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
				if num%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				return isPrime
			}
			return isPrime
		}

		if primeBool(currentNumber) {
			primeCounter++
		}

		if primeCounter == number {
			break
		}

		currentNumber++
	}

	return currentNumber
}

func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
