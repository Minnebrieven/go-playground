package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intToRoman(number int) string {
	maxNumber := 3999
	if number > maxNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func main() {
	var num int
	fmt.Print("Masukan Angka = ")
	fmt.Scan(&num)

	fmt.Println(num, " => ", intToRoman(num))
}
