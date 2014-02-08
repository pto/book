package main

import (
	"fmt"
	"strings"
)

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

	names = " Antônio\tAndré\tFriedrish\t\t\tJean\t\tÉlisabeth\tIsabella \t"
	names = strings.Replace(names, "\t", " ", -1)
	fmt.Printf("|%s|\n", names)
	fmt.Println()

}
