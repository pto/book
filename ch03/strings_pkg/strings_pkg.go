package main

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
	"unicode"
)

const LOOPS = 10000

func main() {
	names := "Niccolò•Noël•Geoffrey•Amélie••Turlough•José"
	fmt.Print("|")
	for _, name := range strings.Split(names, "•") {
		fmt.Printf("%s|", name)
	}
	fmt.Print("\n|")
	for _, name := range strings.SplitAfter(names, "•") {
		fmt.Printf("%s|", name)
	}
	fmt.Println()

	for _, record := range []string{
		"László Lajtha*1892*1963",
		"Édouard Lalo\t1823\t1892",
		"José Ángel Lamas|1775|1814"} {
		fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
			switch char {
			case '\t', '*', '|':
				return true
			}
			return false
		}))
	}
	fmt.Println()

	names = " Antônio\tAndré\tFriedrich\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	names = strings.Replace(names, "\t", " ", -1)
	fmt.Printf("|%s|\n", names)
	fmt.Println(strings.Fields(names))
	fmt.Printf("|%s|\n", strings.Join(strings.Fields(names), " "))
	fmt.Printf("|%s|\n", SimpleSimplifyWhitespace(names))
	fmt.Printf("|%s|\n", SimplifyWhitespace(names))
	fmt.Println()

	start := time.Now()
	for i := 0; i < LOOPS; i++ {
		MySimplifyWhitespace(names)
	}
	fmt.Println("MySimplifyWhitespace duration:", time.Since(start))

	start = time.Now()
	for i := 0; i < LOOPS; i++ {
		SimplifyWhitespace(names)
	}
	fmt.Println("SimplifyWhitespace duration:", time.Since(start))

	start = time.Now()
	for i := 0; i < LOOPS; i++ {
		RegExpSimplifyWhitespace(names)
	}
	fmt.Println("RegExpSimplifyWhitespace duration:", time.Since(start))
	fmt.Println()

	asciiOnly := func(char rune) rune {
		if char > 127 {
			return '?'
		}
		return char
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich 💤  !"))
	fmt.Println()

	reader := strings.NewReader("Café 🚂  ❉")
	for {
		char, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%U '%c' %d: % X\n", char, char, size, []byte(string(char)))
	}
}

func MySimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func SimpleSimplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func SimplifyWhitespace(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}
	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

var simplifyWhitespaceRx = regexp.MustCompile(`[\s\p{Zl}\p{Zp}]+`)

func RegExpSimplifyWhitespace(s string) string {
	return strings.TrimSpace(simplifyWhitespaceRx.ReplaceAllLiteralString(s, " "))
}
