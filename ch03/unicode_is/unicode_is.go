package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <character>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	char, length := utf8.DecodeRuneInString(os.Args[1])
	fmt.Printf("Character %c has UTF8 length of %d\n", char, length)
	fmt.Println("Properties:")
	print_rune_is(char, unicode.Properties)
	fmt.Println("Types:")
	print_rune_types(char)
	fmt.Println("Scripts:")
	print_rune_is(char, unicode.Scripts)
}

func print_rune_is(char rune, props map[string]*unicode.RangeTable) {
	for prop, table := range props {
		if unicode.Is(table, char) {
			fmt.Println("  ", prop)
		}
	}
}

func print_rune_types(char rune) {
	if unicode.IsControl(char) {
		fmt.Println("   Control")
	}
	if unicode.IsDigit(char) {
		fmt.Println("   Digit")
	}
	if unicode.IsGraphic(char) {
		fmt.Println("   Graphic")
	}
	if unicode.IsLetter(char) {
		fmt.Println("   Letter")
	}
	if unicode.IsLower(char) {
		fmt.Println("   Lower")
	}
	if unicode.IsMark(char) {
		fmt.Println("   Mark")
	}
	if unicode.IsPrint(char) {
		fmt.Println("   Print")
	}
	if unicode.IsPunct(char) {
		fmt.Println("   Punct")
	}
	if unicode.IsSpace(char) {
		fmt.Println("   Space")
	}
	if unicode.IsSymbol(char) {
		fmt.Println("   Symbol")
	}
	if unicode.IsTitle(char) {
		fmt.Println("   Title")
	}
	if unicode.IsUpper(char) {
		fmt.Println("   Upper")
	}
}
