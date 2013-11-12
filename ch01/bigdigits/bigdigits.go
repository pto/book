package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" ||
		(strings.HasPrefix(os.Args[1], "-") && os.Args[1] != "-b" &&
			os.Args[1] != "--bar") || (len(os.Args) == 2 &&
		(os.Args[1] == "-b" || os.Args[1] == "--bar")) {
		fmt.Printf("usage: %s [-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	var showBar = false
	var stringOfDigits string
	if os.Args[1] == "-b" || os.Args[1] == "--bar" {
		showBar = true
		stringOfDigits = os.Args[2]
	} else {
		stringOfDigits = os.Args[1]
	}

	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if digit >= 0 && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatalf("%s: invalid whole number",
					filepath.Base(os.Args[0]))
			}
		}
		if showBar && row == 0 {
			fmt.Println(strings.Repeat("*", len(line)-2))
		}
		fmt.Println(line)
		if showBar && row == len(bigDigits[0])-1 {
			fmt.Println(strings.Repeat("*", len(line)-2))
		}
	}
}

var bigDigits = [][]string{{
	" 000 ",
	"0   0",
	"0   0",
	"0   0",
	"0   0",
	"0   0",
	" 000 "}, {

	" 1 ",
	"11 ",
	" 1 ",
	" 1 ",
	" 1 ",
	" 1 ",
	"111"}, {

	" 222 ",
	"2   2",
	"   2 ",
	"  2  ",
	"2    ",
	"2    ",
	"22222"}, {

	" 333 ",
	"3   3",
	"    3",
	"  33 ",
	"    3",
	"3   3",
	" 333 "}, {

	"   4 ",
	"  44 ",
	" 4 4 ",
	"44444",
	"   4 ",
	"   4 ",
	"   4 "}, {

	"55555",
	"5    ",
	"5    ",
	"5555 ",
	"    5",
	"    5",
	"5555 "}, {

	" 666 ",
	"6   6",
	"6    ",
	"6 66 ",
	"66  6",
	"6   6",
	" 666 "}, {

	"77777",
	"7   7",
	"   7 ",
	"  7  ",
	" 7   ",
	"7    ",
	"7    "}, {

	" 888 ",
	"8   8",
	"8   8",
	" 888 ",
	"8   8",
	"8   8",
	" 888 "}, {

	" 999 ",
	"9   9",
	"9  99",
	" 99 9",
	"    9",
	"9   9",
	" 999 "}}
