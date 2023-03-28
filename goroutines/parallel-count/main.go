package main

import (
	"fmt"
	"sync"
)

func splitText(text string) (string, string) {
	middleIndex := len(text) / 2
	return text[:middleIndex], text[middleIndex:]
}

func countString(text string, wg *sync.WaitGroup, mapString map[string]int) {
	defer wg.Done()
	for i := 0; i <= len(text)-1; i++ {
		str := string(text[i])
		if str == " " {
			str = "space"
		}
		mapString[str]++
		fmt.Println(str, ":", mapString[str])
	}

}

func main() {
	mapString := make(map[string]int)
	var wg sync.WaitGroup
	var text string = "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqual"
	text1, text2 := splitText(text)
	text1, text3 := splitText(text1)
	text2, text4 := splitText(text2)

	wg.Add(4)
	go countString(text1, &wg, mapString)
	go countString(text2, &wg, mapString)
	go countString(text3, &wg, mapString)
	go countString(text4, &wg, mapString)

	fmt.Println("Waiting goroutine to finish")
	wg.Wait()
	fmt.Println("Goroutine is finish")
	for key, value := range mapString {
		fmt.Println(key, " => ", value)
	}
}
