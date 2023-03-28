package main

import (
	"fmt"
	"strconv"
)

type Pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []Pair {
	itemCounter := map[string]int{} //map for counting items
	for _, item := range items {
		if itemCounter[item] == 0 {
			itemCounter[item] = 1
		} else {
			itemCounter[item]++
		}
	}

	sortArray := make([][2]string, len(itemCounter))
	func() {
		//convert map to 2D-array for sorting process

		idx := 0
		for key, value := range itemCounter {
			sortArray[idx] = [2]string{key, strconv.Itoa(value)}
			idx++
		}
	}()

	func() {
		//bubble sort
		for {
			var swapped bool
			for i := 0; i < len(sortArray); i++ {
				if i == len(sortArray)-1 {
					continue
				}
				leftElement, rightElement := sortArray[i], sortArray[i+1]
				if leftElement[1] > rightElement[1] {
					sortArray[i] = rightElement
					sortArray[i+1] = leftElement
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}
	}()

	resPair := []Pair{}
	func() {
		//convert sortArray(sorted) to slice of Pair ([]Pair)

		for idx, _ := range sortArray {
			pairCount, err := strconv.Atoi(sortArray[idx][1])
			if err != nil {
				pairCount = itemCounter[sortArray[idx][0]]
			}
			itemPair := Pair{name: sortArray[idx][0], count: pairCount}
			resPair = append(resPair, itemPair)
		}
	}()

	return resPair
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	//golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// //C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// //football->1 basketball->1 tenis->1
}
