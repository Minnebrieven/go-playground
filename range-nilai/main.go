package main

import "fmt"

func rangeNilai(nilai int) string {
	switch {
	case nilai <= 100 && nilai >= 80:
		return "A"
	case nilai <= 79 && nilai >= 65:
		return "B"
	case nilai <= 64 && nilai >= 50:
		return "C"
	case nilai <= 49 && nilai >= 35:
		return "D"
	case nilai <= 34 && nilai >= 0:
		return "E"
	default:
		return "Nilai Invalid"
	}
}

func main() {
	var nilai int

	fmt.Print("==========\nRange Nilai\n==========\n")
	fmt.Print("Masukan Nilai = ")
	fmt.Scanf("%d", &nilai)

	result := rangeNilai(nilai)

	fmt.Printf("%d => %s", nilai, result)
	fmt.Println("")
}
