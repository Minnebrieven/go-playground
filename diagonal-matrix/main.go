package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	var diagonalLeftToRight, diagonalRightToLeft int
	var lenght = len(matrix)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if i == j {
				diagonalLeftToRight += matrix[i][j]
				// fmt.Println("i ~> [", i, " ", j, "]~> ", matrix[i][j])
			}

			if i+j == lenght-1 {
				diagonalRightToLeft += matrix[i][j]
				// fmt.Println("j => [", i, " ", j, "]~> ", matrix[i][j])
			}
		}
	}
	fmt.Println("|", diagonalLeftToRight, " - ", diagonalRightToLeft, "|")
	absoluteValue := math.Abs(float64(diagonalLeftToRight - diagonalRightToLeft))
	fmt.Printf("Nilai Mutlak/Absolute => %v", absoluteValue)
}
