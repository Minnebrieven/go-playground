package main

import "fmt"

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	return fibonacci(number-1) + fibonacci(number-2)
	//fibo(8) + fibo(7)
	//fibo(7) + fibo(6)
	//fibo()
}

func main() {
	fmt.Println(fibonacci(0))  //0
	fmt.Println(fibonacci(2))  //1
	fmt.Println(fibonacci(9))  //34
	fmt.Println(fibonacci(10)) //55
	fmt.Println(fibonacci(12)) //144
}
