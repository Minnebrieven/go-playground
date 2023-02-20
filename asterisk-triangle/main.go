package main

import "fmt"

func main() {
	var rows, space, star int

	fmt.Printf("Masukan jumlah baris : ")
	fmt.Scanf("%d", &rows)
	fmt.Println("")

	for i := 1; i <= rows; i++ {
		star = 0
		for space = 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			star++
			if star == 2*i-1 {
				// fmt.Printf("%d", star)
				break
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
