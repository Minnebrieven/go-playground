package main

import (
	"fmt"
	"strconv"
)

func faktorBilangan(bilangan int) string {
	var faktor string = "1"

	for i := 2; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor += ", " + strconv.Itoa(i)
		} else {
			continue
		}
	}
	return faktor
}

func main() {
	var bilangan int

	fmt.Print("Masukan Bilangan = ")
	fmt.Scanf("%d", &bilangan)

	result := faktorBilangan(bilangan)

	fmt.Printf("Faktor Bilangan dari %d => %v\n", bilangan, result)
}
