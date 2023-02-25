package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	rawArray := append(arrayA, arrayB...)
	compare := make(map[string]bool)
	noDuplicate := []string{}

	for _, name := range rawArray {
		if _, value := compare[name]; !value {
			compare[name] = true
			noDuplicate = append(noDuplicate, name)
		}
	}
	return noDuplicate
}

func main() {
	//Test Case
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan]

	fmt.Println(ArrayMerge([]string{"alissa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alissa", "law"}))
	// ["alissa", "yoshimitsu", "devil jin", "law]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
