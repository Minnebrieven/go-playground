package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func IsPositive(number int) bool {
	if number < 1 {
		fmt.Println("Hanya dapat menerima bilangan positif")
	}
	return true
}

func waitToFinish(limit int) {
	timeToSecond := limit * 3
	time.Sleep(time.Duration(timeToSecond) * time.Second)
}

func kelipatanX(kelipatan, limit int) {
	for i := 1; i <= limit; i++ {
		currentX := i * kelipatan
		fmt.Println(currentX, "at => ", time.Since(start), " second")
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var number, limit int

	fmt.Print("\nMasukan bilangan kelipatan : ")
	fmt.Scan(&number)
	fmt.Print("\nMasukan jumlah bilangan : ")
	fmt.Scan(&limit)

	go kelipatanX(number, limit)
	waitToFinish(limit) //keep kelipatanX() running to finish by limit * 3 (second)
}
