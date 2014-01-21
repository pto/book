package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%t %t\n", true, false)
	fmt.Printf("%d %d\n", intForBool(true), intForBool(false))
	t := parseBool("True")
	f := parseBool("FALSE")
	fmt.Printf("%t %t\n", t, f)
	fmt.Println()

	fmt.Println("|123456" + strings.Repeat("|123456789", 5) +
		strings.Repeat("|1234567", 2) + "|")
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|% 9b|%+7b|%+7b|\n",
		37, 37, 37, 37, 37, -37, 37, -37)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	i := 3931
	fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i, i, i, i, i, i)
	i = 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
}

func intForBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func parseBool(s string) bool {
	val, err := strconv.ParseBool(s)
	if err != nil {
		fmt.Printf("cannot parse \"%s\"\n", s)
	}
	return val
}

func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}
