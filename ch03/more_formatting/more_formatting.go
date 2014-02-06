package main

import (
	"fmt"
	"math"
)

type polar struct{ radius, θ float64 }

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
	fmt.Println()

	s = "Alias ↔︎ Synonym"
	chars := []rune(s)
	bytes := []byte(s)
	fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)
	fmt.Println()

	i := 5
	f := -48.3124
	s = "Tomás Bretón"
	fmt.Printf("|%p → %d|%p → %f|%#p → %s|\n", &i, i, &f, f, &s, s)
	fmt.Println()

	fmt.Println([]float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%%v:   %v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%%#v:  %#v\n", []float64{math.E, math.Pi, math.Phi})
	fmt.Printf("%%.5f: %.5f\n", []float64{math.E, math.Pi, math.Phi})

	fmt.Printf("%%q: %q\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%%v: %v\n", []string{"Software patents", "kill", "innovation"})
	fmt.Printf("%%#v: %#v\n",
		[]string{"Software patents", "kill", "innovation"})
	fmt.Printf("%%17s: %17s\n",
		[]string{"Software patents", "kill", "innovation"})
	fmt.Println()

	fmt.Printf("%%v: %v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%%#v: %#v\n", map[int]string{1: "A", 2: "B", 3: "C", 4: "D"})
	fmt.Printf("%%v: %v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%%#v: %#v\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
	fmt.Printf("%%04b: %04b\n", map[int]int{1: 1, 2: 2, 3: 4, 4: 8})
}
