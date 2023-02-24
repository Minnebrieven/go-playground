package main

import (
	"fmt"
	//"math"
)

// My original Code
func pow(x, n int) int {
	var i, initialNumber int = 1, x
	for i < n {
		x = x * initialNumber
		// fmt.Println(x)
		i++
	}
	return x
}

/*
//My code using math package from golang

func powerOfNumber(x, n int) int {
	res := math.Pow(float64(x), float64(n))

	return int(res)
}
*/

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
