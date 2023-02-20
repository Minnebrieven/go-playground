package main

import (
	"fmt"
)

func main() {
	var theNumber int //angka yang akan dicheck prima atau bukan

	fmt.Print("Masukan Angka/Bilangan Prima : ")
	fmt.Scanf("%d", &theNumber) //masukan angka yang akan dicheck dibaris ini

	fmt.Printf("\nBilangan Prima = true | Bukan Bilangan Prima = false \nAngka %d => %t \n", theNumber, checkBilangan((theNumber)))
}

// bilangan prima adalah bilangan yang bisa dibagi 1 atau bilangan itu sendiri
func checkBilangan(theNum int) (isPrime bool) { //false = bukan bilangan prima || true = bilangan prima
	switch {
	case theNum == 1:
		return
	case theNum == 2 || theNum == 3: //2 dan 3 adalah bilangan prima
		isPrime = true
	case theNum%2 != 0 && theNum%3 != 0: //angka yang tidak bisa dibagi 2 atau 3
		isPrime = true
	}
	return
}
