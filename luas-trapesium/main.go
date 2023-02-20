package main

import "fmt"

func luasTrapesium(sisiA, sisiB, tinggiTrapesium int) int {
	sumAB := sisiA + sisiB
	return sumAB * tinggiTrapesium / 2
}

func main() {
	var sisiA, sisiB, tinggiTrapesium int

	fmt.Println("Masukan panjang sisi A, B dan Tinggi Trapesium. Pisahkan nilai dengan menggunakan spasi (space)")
	fmt.Scanln(&sisiA, &sisiB, &tinggiTrapesium)

	result := luasTrapesium(sisiA, sisiB, tinggiTrapesium)

	fmt.Printf("Luas Trapesium adalah %v \n", result)
}
