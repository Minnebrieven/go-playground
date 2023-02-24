package main

import (
	"fmt"
)

// bilangan prima adalah bilangan yang bisa dibagi 1 atau bilangan itu sendiri
func primeNumber(number int) bool { //false = bukan bilangan prima || true = bilangan prima
	switch {
	case number == 1:
		return false
	case number == 2 || number == 3: //2 dan 3 adalah bilangan prima
		return true
	case number%2 != 0 && number%3 != 0 && number%5 != 0: //angka yang tidak bisa dibagi 2, 3 atau 5
		return true
	}
	return false
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
