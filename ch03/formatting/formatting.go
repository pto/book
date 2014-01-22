package main

import (
	"fmt"
	"math"
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
	fmt.Printf("%d %#04x %U '%c' %U\n", 0x3a6, 934, '\u03a6', '\U000003a6', 934)
	fmt.Println()

	fmt.Println(strings.Repeat("|12345678901234567890", 3) + "|")
	for _, x := range []float64{-.258, 7194.84, -60897162.0218, 1.500089e-8} {
		fmt.Printf("|%20.5e|%20.5f|%s|\n", x, x, Humanize(x, 20, 5, '*', ','))
	}
	fmt.Println()

	for _, x := range []complex128{
		2 + 3i, 172.6 - 58.3019i, -.827e2 + 9.04831e-3i} {
		fmt.Printf("|%15s|%9.3f|%.2f|%.1e|\n",
			fmt.Sprintf("%6.2f%+.3fi", real(x), imag(x)), x, x, x)
	}
	fmt.Println()

	slogan := "End Óréttlæti♥"
	fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)
	chars := []rune(slogan)
	fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)
	bytes := []byte(slogan)
	fmt.Printf("%s\n%x\n%X\n% X\n%v\n% #x\n",
		bytes, bytes, bytes, bytes, bytes, bytes)
	fmt.Println()

	s := "Dare to be naïve"
	fmt.Printf("|%22s|%-22s|%10s|%5s|\n", s, s, s, s)
	i = strings.Index(s, "n")
	fmt.Printf("|%.10s|%.*s|%-22.10s|%s|\n", s, i, s, s, s)
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

func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:]
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:]
	}
	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}
	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
