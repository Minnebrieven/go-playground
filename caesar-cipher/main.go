package main

import "fmt"

var AsciiSetStringToInt = map[string]int{"a": 97, "b": 98, "c": 99, "d": 100, "e": 101, "f": 102, "g": 103, "h": 104, "i": 105, "j": 106, "k": 107, "l": 108, "m": 109, "n": 110, "o": 111, "p": 112, "q": 113, "r": 114, "s": 115, "t": 116, "u": 117, "v": 118, "w": 119, "x": 120, "y": 121, "z": 122}
var AsciiSetIntToString = map[int]string{97: "a", 98: "b", 99: "c", 100: "d", 101: "e", 102: "f", 103: "g", 104: "h", 105: "i", 106: "j", 107: "k", 108: "l", 109: "m", 110: "n", 111: "o", 112: "p", 113: "q", 114: "r", 115: "s", 116: "t", 117: "u", 118: "v", 119: "w", 120: "x", 121: "y", 122: "z"}

type Caesar struct {
	text   string
	offset int
}

func (c Caesar) encodeText() string {
	var resultText string

	// fmt.Println(c.text)
	offset := c.offset
	// fmt.Println("Initial Offset ", offset)
	func() {
		if c.offset > 26 {
			offset = offset % 26
		}
	}()
	// fmt.Println("Final Offset ", offset)

	for i := 0; i < len(c.text); i++ {
		asciiValue := AsciiSetStringToInt[string(c.text[i])]
		// fmt.Print(" Initial Char ", AsciiSetIntToString[asciiValue], " => ", asciiValue)
		asciiValue += offset
		func() {
			if asciiValue > 122 {
				asciiValue -= 26
			}
		}()
		// fmt.Print(" Final Char ", AsciiSetIntToString[asciiValue], " => ", asciiValue, "\n")
		resultText += AsciiSetIntToString[asciiValue]
	}
	return resultText
}

func caesar(offset int, input string) string {
	ces := Caesar{
		text:   input,
		offset: offset,
	}
	return ces.encodeText()
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
