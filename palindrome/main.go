package main

import (
	"fmt"
)

func palindromeChecker(word string) bool {
	for i := 0; i <= (len(word) / 2); i++ {
		j := len(word) - 1 - i
		if word[i] != word[j] {
			return false
		}
	}
	return true
}

func main() {
	var theWord string

	fmt.Print("Palindrome word Checker. Input your word down here: ")
	fmt.Scanln("%s", &theWord)

	fmt.Printf("\nIs your word (%s) palindrome ? %v", theWord, palindromeChecker(theWord))
}
