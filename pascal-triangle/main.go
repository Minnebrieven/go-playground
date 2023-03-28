package main

import "fmt"

func generatePascal(rows int) [][]int {
	var tempNumber int
	pascalTriangle := [][]int{}
	for i := 0; i < rows; i++ {
		tempCols := []int{}
		switch {
		case i == 0:
			tempCols = append(tempCols, 1)
		case i == 1:
			tempCols = append(tempCols, 1, 1)
		default:
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					tempNumber = 1
					tempCols = append(tempCols, tempNumber)
					continue
				}
				tempNumber = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
				tempCols = append(tempCols, tempNumber)
			}
		}
		pascalTriangle = append(pascalTriangle, tempCols)
	}
	return pascalTriangle
}

func main() {
	var rows int
	fmt.Print("Masukan Tinggi/Baris dari segitiga pascal : ")
	fmt.Scan(&rows)

	fmt.Println(generatePascal(rows))
}
