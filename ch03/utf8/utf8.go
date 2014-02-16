package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "¶ Greetings!"
	r, l := utf8.DecodeRuneInString(s)
	l2 := utf8.RuneLen(r)
	ok := utf8.ValidString(s)
	fmt.Printf("rune %c length %d = %d ok %t\n", r, l, l2, ok)
}
