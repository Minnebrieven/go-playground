package main

import (
	"fmt"
	"time"
)

func main() {
	var kelipatan, limit int = 3, 10
	cNumber := make(chan int, limit)

	go func() {
		for i := 1; i <= limit; i++ {
			currentX := i * kelipatan
			cNumber <- currentX
		}
	}()

	go func() {
		for i := 0; i < limit; i++ {
			foundNumber := <-cNumber
			fmt.Println(i, "=> kelipatan ", kelipatan, " => ", foundNumber)
		}
	}()
	<-time.After(time.Second * 5)
}
