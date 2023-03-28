package main

import (
	"fmt"
	"math"
)

func Frog(jump []int) int {
	n := len(jump)
	cost := make([]int, n)
	for i := range cost {
		cost[i] = math.MaxInt32
	}

	cost[0] = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j <= i+2 && j < n; j++ {
			cost[j] = min(cost[j], cost[i]+abs(jump[i]-jump[j]))
		}
	}

	return cost[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
