package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	countAngka := make(map[int]int)
	onceAppear := []int{}
	i := 0
	for i < len(angka) {
		parsedAngka, _ := strconv.Atoi(angka[i : i+1])
		countAngka[parsedAngka] += 1
		i++
	}
	for key, val := range countAngka {
		if val == 1 {
			onceAppear = append(onceAppear, key)
		}
	}
	return onceAppear
}

func main() {
	fmt.Println(munculSekali("1234123"))    //[4]
	fmt.Println(munculSekali("76523752"))   //[6,3]
	fmt.Println(munculSekali("12345"))      //[1,2,3,4,5]
	fmt.Println(munculSekali("1122334455")) //[]
	fmt.Println(munculSekali("0872504"))    //[8, 7, 2, 5, 4]
}
