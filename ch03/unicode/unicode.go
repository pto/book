package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("\u212bngstrom")
	fmt.Println("\u00c5 or A\u030a or \x41\xcc\x8a")
	fmt.Println("\u212b" == "\xe2\x84\xab")
	fmt.Println("A\u030a" == "\x41\xcc\x8a")
	fmt.Printf("% [1]x %[1]v\n", "A\u030a")
	fmt.Printf("% [1]x %[1]v\n", "\x41\xcc\x8a")
	fmt.Println()

	æs := ""
	for _, char := range []rune{'\u00e6', 0xe6, 0346, 230, '\xE6', '\u00e6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æs += string(char)
	}
	fmt.Println(æs)
	fmt.Println()

	fmt.Println(IsHexDigit('8'), IsHexDigit('x'), IsHexDigit('X'),
		IsHexDigit('b'), IsHexDigit('B'))
}

func IsHexDigit(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}
