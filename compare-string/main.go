package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) < len(b) {
		txtIndex := strings.Index(b, a)
		text := b[txtIndex : txtIndex+len(a)]
		return text
	} else {
		txtIndex := strings.Index(a, b)
		text := a[txtIndex : txtIndex+len(b)]
		return text
	}

}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //	AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //	KANG
	fmt.Println(Compare("KI", "KIJANG"))      //	KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //	KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //	ILA
}
