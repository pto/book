package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	phrase := "naïve\u2028🚀"
	fmt.Printf("string: %s\n", phrase)
	fmt.Println("len:", len(phrase),
		"RuneCountInString:", utf8.RuneCountInString(phrase))
	fmt.Println("index rune     char bytes")
	for index, char := range phrase {
		fmt.Printf(" %2d   %-8U  %c    % X\n",
			index, char, char, []byte(string(char)))
	}
	last_rune, length := utf8.DecodeLastRuneInString(phrase)
	fmt.Printf("last rune: 0x%x length: %d\n", last_rune, length)
	fmt.Printf("length of string: %d length of []rune: %d\n",
		len(phrase), len([]rune(phrase)))
	fmt.Printf("line separator: %[1]U 0x%[1]x '%[1]c' %s\n", '\u2028', string('\u2028'))
	fmt.Printf("paragraph separator: %[1]U 0x%[1]x '%[1]c' %s\n", '\u2029', string('\u2029'))

	line := "rå tørt\u2028vær"
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line[j:])
	lastWord := line[j+size:]
	fmt.Println("size of space:", size)
	fmt.Println(firstWord, lastWord)
}
