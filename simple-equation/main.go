package main

import (
	"fmt"
	"math"
)

func simpleEquations(a, b, c int) {
	var abcRange int = 10000
	var solutionContainer []int

	defer func() {
		if len(solutionContainer) == 0 {
			fmt.Println("No solution found")
		} else {
			fmt.Println(solutionContainer[0], solutionContainer[1], solutionContainer[2])
		}
	}()

	func() {
		max := math.Max(float64(a), float64(b))
		max = math.Max(max, float64(c))

		if int(max) < 100 {
			abcRange = 1000
		}
	}()

	xyzA := func(x, y, z, a int) bool {
		return x+y+z == a
	}
	xyzB := func(x, y, z, B int) bool {
		return x*y*z == B
	}
	xyzC := func(x, y, z, C int) bool {
		return x*x+y*y+z*z == C
	}

	for x := 1; x <= abcRange; x++ {
		for y := 1; y <= abcRange; y++ {
			for z := 1; z < abcRange; z++ {
				if xyzA(x, y, z, a) && xyzB(x, y, z, b) && xyzC(x, y, z, c) {
					solutionContainer = append(solutionContainer, x, y, z)
				}
				if len(solutionContainer) != 0 {
					break
				}
			}
			if len(solutionContainer) != 0 {
				break
			}
		}
		if len(solutionContainer) != 0 {
			break
		}
	}
}

func main() {
	simpleEquations(1, 2, 3)
	simpleEquations(6, 6, 14)
}
