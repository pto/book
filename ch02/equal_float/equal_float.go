package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("math.SmallestNonzeroFloat64 is", math.SmallestNonzeroFloat64)
	fmt.Println()

	x, y := 0.0, 0.0
	fmt.Printf("%-5s %-5s %-5s %-5s %-5s %-5s ", "==", "(-1)", "e-15", "e-16", "Pr6", "Pr60")
	fmt.Println("x, y")
	for i := 0; i < 10; i++ {
		x += 0.1
		if i%2 == 0 {
			y += 0.2
		} else {
			fmt.Printf("%-5t %-5t %-5t %-5t %-5t %-5t ", x == y, EqualFloat(x, y, -1),
				EqualFloat(x, y, 1e-15), EqualFloat(x, y, 1e-16), EqualFloatPrec(x, y, 6),
				EqualFloat(x, y, 60))
			fmt.Printf("%v, %v\n", x, y)
		}
	}
}

func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

func EqualFloatPrec(x, y float64, decimals int) bool {
	a := fmt.Sprintf("%.*f", decimals, x)
	b := fmt.Sprintf("%.*f", decimals, y)
	return len(a) == len(b) && a == b
}
