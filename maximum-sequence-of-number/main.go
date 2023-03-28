package main

import "fmt"

func MaxSequence(arr []int) int {
	var tempSum, initialIdx, lastIdx int = 0, 0, len(arr) - 1

	//skipping below zero value at the start of equence
	func() {
		for i := 0; i < len(arr); i++ {
			if arr[i] > 0 {
				initialIdx = i
				break
			}
		}
	}()

	//skipping below zero value at the end of sequence
	func() {
		for i := len(arr) - 1 - 1; i == 0; i-- {
			if arr[i] > 0 {
				lastIdx = i
				break
			}
		}
	}()

	startSequence := arr[initialIdx:lastIdx]
	maxSum := startSequence[0]
	for i := 0; i < len(startSequence); i++ {
		if maxSum < tempSum {
			maxSum = tempSum
		}

		tempSum += startSequence[i]
		if tempSum < -1 {
			tempSum = 0
			continue
		}
	}

	switch {
	case tempSum < maxSum:
		return maxSum
	default:
		return tempSum
	}
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
