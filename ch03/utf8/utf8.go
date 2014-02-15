package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "¶ Greetings!"
	r, l := utf8.DecodeRuneInString(s)
	fmt.Printf("rune %c length %d\n", r, l)
}
