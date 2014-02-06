package main

import (
	"fmt"
	"math"
)

type polar struct{ radius, theta float64 }

func main() {
	p := polar{-83.40, 71.60}
	fmt.Printf("|%%T|%%v|%%#v|\n")
	fmt.Printf("|%T|%v|%#v|\n", p, p, p)
	fmt.Printf("|%T|%v|%#v|\n", false, false, false)
	fmt.Printf("|%T|%v|%#v|\n", 7607, 7607, 7607)
	fmt.Printf("|%T|%v|%#v|\n", math.E, math.E, math.E)
	fmt.Printf("|%T|%v|%#v|\n", 5+7i, 5+7i, 5+7i)

	s := "Relativity"
	fmt.Printf("|%%T|\"%%v\"|\"%%s\"|%%q|\n")
	fmt.Printf("|%T|\"%v\"|\"%s\"|%q|\n", s, s, s, s)
}
