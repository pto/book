package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, truth := range []string{"1", "t", "TRUE", "false", "F", "0", "5"} {
		if b, err := strconv.ParseBool(truth); err != nil {
			fmt.Printf("\n%v", err)
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Println()

	x, err := strconv.ParseFloat("-99.7", 64)
	fmt.Printf("%8T %6v %v\n", x, x, err)
	y, err := strconv.ParseInt("71309", 10, 16)
	fmt.Printf("%8T %6v %v\n", y, y, err)
	z, err := strconv.Atoi("71309")
	fmt.Printf("%8T %6v %v\n", z, z, err)
	fmt.Println()

	s := strconv.FormatBool(z > 100)
	fmt.Println(s)
	i, err := strconv.ParseInt("0xDEED", 0, 32)
	fmt.Println(i, err)
	j, err := strconv.ParseInt("0707", 0, 32)
	fmt.Println(j, err)
	k, err := strconv.ParseInt("10111010001", 2, 32)
	fmt.Println(k, err)

	m := 16769023
	fmt.Println(strconv.Itoa(m))
	fmt.Println(strconv.FormatInt(int64(m), 10))
	fmt.Println(strconv.FormatInt(int64(m), 2))
	fmt.Println(strconv.FormatInt(int64(m), 16))
	fmt.Println()

	s = "Alle ønsker å være fri."
	quoted := strconv.QuoteToASCII(s)
	fmt.Println(quoted)
	fmt.Println(strconv.Unquote(quoted))

	var bs []byte
	bs = strconv.AppendQuote(bs, "👍")
	fmt.Printf("|%s|% #x|%d|\n", string(bs), bs, bs)
}
