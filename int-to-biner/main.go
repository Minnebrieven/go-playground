package main

import (
	"fmt"
	"strconv"
)

func intToBiner(n int) []string {
	var binarySlice []string

	for i := 0; i <= n; i++ {
		toBinary := strconv.FormatInt(int64(i), 2)
		binarySlice = append(binarySlice, toBinary)
	}

	return binarySlice
}

func main() {
	var n int
	fmt.Print("Masukkan angka yang akan dikonversi: ")
	fmt.Scanln(&n)

	fmt.Println(intToBiner(n))
}
