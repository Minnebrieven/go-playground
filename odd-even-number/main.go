package main

import "fmt"

func oddEvenChecker(number int) string {
	if number%2 == 0 {
		return "Genap / Even"
	} else {
		return "Ganjil / Odd"
	}
}

func main() {
	var Number int

	fmt.Print("Masukan Angka = ")
	fmt.Scanf("%d", &Number)

	result := oddEvenChecker(Number)

	fmt.Printf("Angka %d => %s \n", Number, result)
}
