package main

import "fmt"

func Mapping(slice []string) map[string]int {
	mapString := make(map[string]int)

	for _, str := range slice {
		if mapString[str] == 0 {
			mapString[str] = 1
		} else if mapString[str] != 0 {
			mapString[str]++
		}
	}
	return mapString
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}
